package setter

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pb/setter"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[setter.FileOptions](file.Desc.Options(), setter.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.setter"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "setter", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()

		// https://github.com/protocolbuffers/protobuf-go/blob/v1.28.1/cmd/protoc-gen-go/internal_gengo/main.go
		// Blank imports are automatically handled by g.Import.
		// Packages referenced by g.QualifiedGoIdent are automatically imported.
		// See documents for g.Import.
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
		//	g.Import(f.GoImportPath)
		//}

		generate(g, file.Messages)
	}
	return nil
}

func generate(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		generate(g, m.Messages)

		messageOpts := pkg.ProtoGetExtension[setter.MessageOptions](m.Desc.Options(), setter.E_MessageOptions)
		if !messageOpts.GetEnable() {
			continue
		}

		for _, f := range m.Fields {
			fieldName := f.GoName
			fieldOpts := pkg.ProtoGetExtension[setter.FieldOptions](f.Desc.Options(), setter.E_FieldOptions)
			name := strings.TrimSpace(fieldOpts.GetName())
			if name == "" {
				name = fieldName
			}

			typ, _, _ := pkg.GoFieldTypeInfo(g, f)

			g.P("// Setter for field ", fieldName, ".")
			g.P("func (x *", m.GoIdent.GoName, ") Set", name, "(v ", typ, ") {")
			g.P("if x != nil {")
			g.P("x.", fieldName, "=v")
			g.P("}")
			g.P("}")
			g.P()
		}
	}
}
