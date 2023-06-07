package setter

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

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
		for i := 0; i < file.Desc.Imports().Len(); i++ {
			imp := file.Desc.Imports().Get(i)
			f, ok := plugin.FilesByPath[imp.Path()]
			if !ok {
				continue
			}
			// Do not import self.
			if f.GoImportPath == file.GoImportPath || imp.IsWeak {
				continue
			}
			g.Import(f.GoImportPath)
		}

		for _, message := range file.Messages {
			messageOpts := pkg.ProtoGetExtension[setter.MessageOptions](message.Desc.Options(), setter.E_MessageOptions)
			if messageOpts.GetDisable() {
				continue
			}

			for _, field := range message.Fields {
				fieldOpts := pkg.ProtoGetExtension[setter.FieldOptions](field.Desc.Options(), setter.E_FieldOptions)
				if fieldOpts.GetDisable() {
					continue
				}

				g.P("func (x *", message.Desc.Name(), ") Set", field.GoName, "(v ", fieldType(g, field), ") {")
				g.P("if x != nil {")
				g.P("x.", field.GoName, "=v")
				g.P("}")
				g.P("}")
				g.P()
			}
		}
	}
	return nil
}

func fieldType(g *protogen.GeneratedFile, field *protogen.Field) string {
	var (
		typ  string // identifier
		star string // pointer
	)
	if field.Desc.HasPresence() {
		star = "*"
	}

	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		typ = "bool"
	case protoreflect.EnumKind:
		typ = g.QualifiedGoIdent(field.Enum.GoIdent)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		typ = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		typ = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		typ = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		typ = "uint64"
	case protoreflect.FloatKind:
		typ = "float32"
	case protoreflect.DoubleKind:
		typ = "float64"
	case protoreflect.StringKind:
		typ = "string"
	case protoreflect.BytesKind:
		typ = "[]byte"
		star = ""
	case protoreflect.MessageKind, protoreflect.GroupKind:
		typ = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
		star = ""
	}

	if field.Desc.IsList() {
		return "[]" + typ
	} else if field.Desc.IsMap() {
		k := fieldType(g, field.Message.Fields[0])
		v := fieldType(g, field.Message.Fields[1])
		return fmt.Sprintf(`map[%s]%s`, k, v)
	}

	return star + typ
}
