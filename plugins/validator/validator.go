package validator

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/guader/protogen/pb/validator"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[validator.FileOptions](file.Desc.Options(), validator.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.validator"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "validator", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()

		g.P("import (")
		g.P(`"errors"`)
		g.P(`"unicode/utf8"`)
		g.P(")")
		g.P()
		g.P("// Ensure imports are used.")
		g.P("var (")
		g.P("_ = errors.New")
		g.P("_ = utf8.RuneCountInString")
		g.P(")")

		for _, message := range file.Messages {
			g.P("func (x *", message.Desc.Name(), ") Validate() error {")
			g.P("if x == nil {")
			g.P("return nil")
			g.P("}")
			g.P()

			for _, field := range message.Fields {
				opts := pkg.ProtoGetExtension[validator.FieldRules](field.Desc.Options(), validator.E_Rules)
				if opts == nil {
					continue
				}

				isList := field.Desc.IsList()
				isOptional := field.Desc.HasOptionalKeyword()

				g.P(fmt.Sprintf("// Field %s.%s.", field.Parent.Desc.Name(), field.GoName))
				g.P("{")

				if isList {
					for _, code := range generateRepeatedRules(field, opts.GetRepeated()) {
						g.P(code)
					}
					g.P("// Repeated.")
					g.P("for _, v := range x.Get", field.GoName, "() {")
					if field.Desc.Kind() == protoreflect.MessageKind {
						if opts.GetRepeated().GetSkipNil() {
							g.P("// Optional.")
							g.P("if v == nil { continue }")
						} else {
							msg := fmt.Sprintf("%s.%s required.", field.Parent.Desc.Name(), field.GoName)
							g.P("// Required.")
							g.P("if v == nil {")
							g.P(fmt.Sprintf("return errors.New(%#v)", msg))
							g.P("}")
						}
					}
				} else {
					if isOptional {
						g.P("// Optional.")
						g.P("if x.", field.GoName, "!= nil {")
					} else {
						// Can not be nil if is ptr.
						if field.Desc.HasPresence() {
							msg := fmt.Sprintf("%s.%s required.", field.Parent.Desc.Name(), field.GoName)
							g.P("// Required.")
							g.P("if x.", field.GoName, "== nil {")
							g.P(fmt.Sprintf("return errors.New(%#v)", msg))
							g.P("}")
						}
					}
					g.P("v := x.Get", field.GoName, "()")
				}

				ruleCodes := generateRules(field, opts)
				if len(ruleCodes) > 0 {
					for _, code := range ruleCodes {
						g.P(code)
					}
				} else {
					g.P("_ = v")
				}

				if isOptional {
					g.P("}")
				}
				if isList {
					g.P("}")
				}

				g.P("}")
				g.P()
			}

			g.P("return nil")
			g.P("}")
			g.P()
		}
	}

	return nil
}

func generateRules(f *protogen.Field, opts *validator.FieldRules) []string {
	var codes []string
	switch opts.GetType().(type) {
	case *validator.FieldRules_Bool:
		r := opts.GetBool()
		codes = append(codes, generateOperatorRules(f, r.Eq, nil, nil, nil, nil, nil)...)
	case *validator.FieldRules_Enum:
		r := opts.GetEnum()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateEnumRules(f, r)...)
	case *validator.FieldRules_Int32:
		r := opts.GetInt32()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_Uint32:
		r := opts.GetUint32()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_Int64:
		r := opts.GetInt64()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_Uint64:
		r := opts.GetUint64()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_Float:
		r := opts.GetFloat()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_Double:
		r := opts.GetDouble()
		codes = append(codes, generateOperatorRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_String_:
		r := opts.GetString_()
		codes = append(codes, generateStringRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
		codes = append(codes, generateExistenceRules(f, r.In, r.NotIn)...)
	case *validator.FieldRules_Bytes:
		r := opts.GetBytes()
		codes = append(codes, generateBytesRules(f, r.Eq, r.Ne, r.Lt, r.Gt, r.Le, r.Ge)...)
	case *validator.FieldRules_Message:
		codes = append(codes, generateMessageRules(f))
	}
	return codes
}
