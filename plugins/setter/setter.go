package setter

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/guader/protogen/config"
	"github.com/guader/protogen/version"
)

func Generate(conf *config.Config, plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	if !conf.Setter.Enabled {
		return nil
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s.%s.go", file.GeneratedFilenamePrefix, conf.Setter.Suffix),
			file.GoImportPath)
		g.P("// Code generated by protoc-gen-go-ext. DO NOT EDIT.")
		g.P("// version: ", version.Version)
		g.P("// source: ", file.Desc.Path())
		g.P()
		g.P("package ", file.GoPackageName)
		g.P()

		for _, message := range file.Messages {
			for _, field := range message.Fields {
				var typ string
				switch field.Desc.Kind() {
				case protoreflect.BoolKind:
					typ = "bool"
				case protoreflect.EnumKind:
					continue
				case protoreflect.Int32Kind, protoreflect.Sint32Kind:
					typ = "int32"
				case protoreflect.Uint32Kind:
					typ = "uint32"
				case protoreflect.Int64Kind, protoreflect.Sint64Kind:
					typ = "int64"
				case protoreflect.Uint64Kind:
					typ = "uint64"
				case protoreflect.Sfixed32Kind:
					continue
				case protoreflect.Fixed32Kind:
					continue
				case protoreflect.FloatKind:
					typ = "float32"
				case protoreflect.Sfixed64Kind:
					continue
				case protoreflect.Fixed64Kind:
					continue
				case protoreflect.DoubleKind:
					typ = "float64"
				case protoreflect.StringKind:
					typ = "string"
				case protoreflect.BytesKind:
					typ = "[]byte"
				case protoreflect.MessageKind:
					continue
				case protoreflect.GroupKind:
					continue
				}

				var star, list string
				if field.Desc.HasPresence() {
					star = "*"
				}
				if field.Desc.IsList() {
					list = "[]"
				}

				g.P("func (x *", message.Desc.Name(), ")Set", field.GoName, "(v ", list, star, typ, "){")
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
