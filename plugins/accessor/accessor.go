package accessor

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pb/accessor"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[accessor.FileOptions](file.Desc.Options(), accessor.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.accessor"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "accessor", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()

		generate(g, file.Messages)
	}
	return nil
}

func generate(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		generate(g, m.Messages)

		messageOpts := pkg.ProtoGetExtension[accessor.MessageOptions](m.Desc.Options(), accessor.E_MessageOptions)
		if !messageOpts.GetEnable() {
			continue
		}

		msgName := m.GoIdent.GoName
		accName := msgName + "Accessor"

		g.P("func (x *", msgName, ") Accesscor() *", accName, "{return &", accName, "{X:x}", "}")
		g.P()

		g.P("type ", accName, " struct { X *", msgName, "}")
		g.P()

		for _, f := range m.Fields {
			fieldName := f.GoName
			fieldOpts := pkg.ProtoGetExtension[accessor.FieldOptions](f.Desc.Options(), accessor.E_FieldOptions)
			name := strings.TrimSpace(fieldOpts.GetName())
			if name == "" {
				name = fieldName
			}
			getterName := strings.TrimSpace(fieldOpts.GetGetter())
			if getterName == "" {
				getterName = name
			}
			setterName := strings.TrimSpace(fieldOpts.GetSetter())
			if setterName == "" {
				setterName = name
			}

			// Get field type information.
			typ, defaultValue, hasStar := pkg.GoFieldTypeInfo(g, f)

			// Generate getter.
			g.P("// Getter for field ", fieldName, ".")
			g.P("func (a *", accName, ") Get", getterName, "()", typ, "{")
			g.P("if a.X != nil {")
			g.P("return a.X.", fieldName)
			g.P("}")
			if hasStar {
				g.P("return nil")
			} else {
				g.P("return ", defaultValue)
			}
			g.P("}")
			g.P()

			// Generate setter.
			g.P("// Setter for field ", fieldName, ".")
			g.P("func (a *", accName, ") Set", setterName, "(v ", typ, ") {")
			g.P("if a.X != nil {")
			g.P("a.X.", fieldName, "= v")
			g.P("}")
			g.P("}")
			g.P()
		}
	}
}
