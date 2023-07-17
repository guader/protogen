package getter

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pb/getter"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[getter.FileOptions](file.Desc.Options(), getter.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.getter"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "getter", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()

		// Generate imports.
		// See plugins/setter/setter.go
		//for i := 0; i < file.Desc.Imports().Len(); i++ {
		//	imp := file.Desc.Imports().Get(i)
		//	f, ok := plugin.FilesByPath[imp.Path()]
		//	if !ok {
		//		continue
		//	}
		//	// Do not import self.
		//	if f.GoImportPath == file.GoImportPath || imp.IsWeak {
		//		continue
		//	}
		//	fmt.Fprintln(os.Stderr, imp)
		//	g.Import(f.GoImportPath)
		//}

		if err := generate(g, file.Messages); err != nil {
			return errors.Join(errors.New(fmt.Sprintf("err on file: %s", file.Desc.Path())), err)
		}
	}
	return nil
}

func generate(g *protogen.GeneratedFile, ms []*protogen.Message) error {
	for _, m := range ms {
		if err := generate(g, m.Messages); err != nil {
			return err
		}

		fieldNameSet := make(map[string]struct{})
		for _, f := range m.Fields {
			fieldNameSet[f.GoName] = struct{}{}
		}
		generatedFieldNameSet := make(map[string]struct{})

		for _, f := range m.Fields {
			fieldOpts := pkg.ProtoGetExtension[getter.FieldOptions](f.Desc.Options(), getter.E_FieldOptions)
			name := strings.TrimSpace(fieldOpts.GetName())
			if len(name) == 0 {
				continue
			}

			if _, ok := fieldNameSet[name]; ok {
				return errors.New(fmt.Sprintf("name conflict with default getter: %s, message: %s, field: %s",
					name, m.GoIdent.GoName, f.GoName))
			}
			if _, ok := generatedFieldNameSet[name]; ok {
				return errors.New(fmt.Sprintf("name conflict with custom getter: %s, message: %s, field: %s",
					name, m.GoIdent.GoName, f.GoName))
			}
			generatedFieldNameSet[name] = struct{}{}

			typ, dft, hasStar := pkg.GoFieldTypeInfo(g, f)
			g.P("func (x *", m.GoIdent.GoName, ") Get", name, "() ", typ, " {")
			g.P("if x != nil {")
			g.P("return x.", f.GoName)
			g.P("}")
			if hasStar {
				g.P("return nil")
			} else {
				g.P("return ", dft)
			}
			g.P("}")
			g.P()
		}
	}
	return nil
}
