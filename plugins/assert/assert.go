package assert

import (
	"fmt"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/guader/protogen/pb/assert"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	commonGeneratedPkgs := make(map[protogen.GoImportPath]struct{})
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[assert.FileOptions](file.Desc.Options(), assert.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.assert"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "assert", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()
		g.P("import (")
		g.P(`"fmt"`)
		g.P(`"unicode/utf8"`)
		g.P(")")
		g.P()
		g.P("var (")
		g.P("_ = fmt.Errorf")
		g.P("_ = utf8.RuneCountInString")
		g.P(")")
		g.P()

		if _, ok := commonGeneratedPkgs[file.GoImportPath]; !ok {
			cprefix := file.GeneratedFilenamePrefix
			cprefix = filepath.Dir(cprefix)
			cprefix = filepath.Join(cprefix, string(file.GoPackageName))

			csuffix := suffix + ".common"

			cg := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", cprefix, csuffix), file.GoImportPath)
			cg.P(pkg.RenderPackageComments(version.Version, "assert", string(file.GoPackageName)))
			cg.P("package ", file.GoPackageName)
			cg.P()

			generateCommon(cg)
			commonGeneratedPkgs[file.GoImportPath] = struct{}{}
		}

		generateMessages(g, file.Messages)
	}
	return nil
}

func generateMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		messageOpts := pkg.ProtoGetExtension[assert.MessageOptions](m.Desc.Options(), assert.E_MessageOptions)
		if !messageOpts.GetEnable() {
			continue
		}

		for _, f := range m.Fields {
			switch {
			case f.Desc.IsList():
				generateSlice(g, m, f)
			case f.Desc.IsMap():
				// TODO
			default:
				switch f.Desc.Kind() {
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind,
					protoreflect.Uint32Kind, protoreflect.Fixed32Kind,
					protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind,
					protoreflect.Uint64Kind, protoreflect.Fixed64Kind,
					protoreflect.FloatKind, protoreflect.DoubleKind:
					generateNumber(g, m, f)
				case protoreflect.StringKind:
					generateString(g, m, f)
				case protoreflect.EnumKind:
					generateEnumerate(g, m, f)
				case protoreflect.BytesKind:
					generateBytes(g, m, f)
				}
			}

			// nillable
			if pkg.FieldGoTypeNillable(f) {
				generateNillable(g, m, f)
			}
		}

		generateMessages(g, m.Messages)
	}
}
