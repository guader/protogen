package enums

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pb/enums"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[enums.FileOptions](file.Desc.Options(), enums.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.enums"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "enums", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()
		g.P(`import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)`)
		g.P()
		g.P(`var (
	_ = codes.Code(0)
	_ = status.Status{}
)`)
		g.P()

		generateEnums(g, file.Enums)
		generateMessages(g, file.Messages)
	}
	return nil
}

func generateEnums(g *protogen.GeneratedFile, es []*protogen.Enum) {
	for _, e := range es {
		typ := e.GoIdent.GoName

		// Generate enum slice.
		g.P("var ", typ, "_enums = []", typ, "{")
		for _, v := range e.Values {
			g.P(v.GoIdent.GoName, ",")
		}
		g.P("}")
		g.P()

		// Generate validation function.
		g.P("func (x ", typ, ") IsValid() bool {")
		g.P("_, ok := ", typ, "_name[int32(x)]")
		g.P("return ok")
		g.P("}")
		g.P()

		// Generate methods.
		enumOpts := pkg.ProtoGetExtension[enums.EnumOptions](e.Desc.Options(), enums.E_EnumOptions)
		if enumOpts.GetErrMethod().GetEnable() {
			names := make([]string, 0, len(e.Values))
			msgs := make([]string, 0, len(e.Values))
			for _, v := range e.Values {
				opts := pkg.ProtoGetExtension[enums.EnumValueOptions](v.Desc.Options(), enums.E_Enum)
				if opts != nil {
					names = append(names, v.GoIdent.GoName)
					msgs = append(msgs, opts.GetMsg())
				}
			}

			g.P("func (x ", typ, ") Error() string {")
			if len(names) > 0 {
				g.P("switch x {")
				for i, name := range names {
					g.P("case ", name, ":")
					g.P("return ", strconv.Quote(msgs[i]))
				}
				g.P("}")
			}
			g.P("return x.String()")
			g.P("}")
			g.P()

			if !enumOpts.GetErrMethod().GetWithGrpc() {
				continue
			}
			g.P("// GRPCStatus implements the interface defined in status.FromError.")
			g.P("func (x ", typ, ") GRPCStatus() *status.Status {")
			g.P("return status.New(codes.Code(x), x.Error())")
			g.P("}")
			g.P()
			g.P("func (x ", typ, ") GRPCError(msg string) error {")
			g.P("return status.Error(codes.Code(x), msg)")
			g.P("}")
			g.P()
			g.P("func (x ", typ, ") GRPCErrorf(format string, a ...any) error {")
			g.P("return status.Errorf(codes.Code(x), format, a...)")
			g.P("}")
			g.P()
		}
	}
}

func generateMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		generateEnums(g, m.Enums)
		generateMessages(g, m.Messages)
	}
}
