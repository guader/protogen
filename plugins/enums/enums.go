package enums

import (
	"fmt"
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

		// Generate validate function.
		g.P("func (x ", typ, ") IsValid() bool {")
		g.P("_, ok := ", typ, "_name[int32(x)]")
		g.P("return ok")
		g.P("}")
		g.P()

		//enumOpts := pkg.ProtoGetExtension[enums.EnumOptions](e.Desc.Options(), enums.E_EnumOptions)

		// Generate parser functions.
		//number := enumOpts.GetParser().GetDefault()
		//value := e.Values[0]
		//for _, v := range e.Values {
		//	if int32(v.Desc.Number()) == number {
		//		value = v
		//		break
		//	}
		//}
		//number = int32(value.Desc.Number())
		//
		//g.P("func (x *", typ, ") FromString(s string) {")
		//g.P("if i, ok := ", typ, "_value[s]; ok {")
		//g.P("*x = ", typ, "(i)")
		//g.P("} else {")
		//g.P("// Default value: ", number, ".")
		//g.P("*x = ", value.GoIdent.GoName)
		//g.P("}")
		//g.P("}")
		//g.P()
	}
}

func generateMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		generateEnums(g, m.Enums)
		generateMessages(g, m.Messages)
	}
}
