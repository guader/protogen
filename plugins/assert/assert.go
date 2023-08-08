package assert

import (
	"fmt"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/guader/protogen/pb/assert"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
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
			if f.Desc.IsList() || f.Desc.IsMap() {
				continue
			}

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
			}
		}

		generateMessages(g, m.Messages)
	}
}

func generateNumber(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	typ, _ := pkg.FieldGoType(g, f)
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	// range
	g.P("// open: 0 for both close, 1 for left open only, 2 for right open only, 3 for both open.")
	g.P("func (x *", mname, ") Assert", fname, "Range(min, max *", typ, ", open byte) error {")
	g.P("v := x.Get", fname, "()")

	g.P("if min != nil {")
	g.P("if open&1 == 0 {")
	g.P("// left close")
	g.P("if v < *min {")
	g.P(`return fmt.Errorf("`, fullName, ` must be greater or equal than %v, value: %v", *min, v)`)
	g.P("}")
	g.P("} else {")
	g.P("// left open")
	g.P("if v <= *min {")
	g.P(`return fmt.Errorf("`, fullName, ` must be greater than %v, value: %v", *min, v)`)
	g.P("}")
	g.P("}")
	g.P("}")

	g.P("if max != nil {")
	g.P("if open&2 == 0 {")
	g.P("// right close")
	g.P("if v > *max {")
	g.P(`return fmt.Errorf("`, fullName, ` must be less or equal than %v, value: %v", *max, v)`)
	g.P("}")
	g.P("} else {")
	g.P("// right open")
	g.P("if v >= *max {")
	g.P(`return fmt.Errorf("`, fullName, ` must be less than %v, value: %v", *max, v)`)
	g.P("}")
	g.P("}")
	g.P("}")

	g.P("return nil")
	g.P("}")
	g.P()

	// greater
	g.P("func (x *", mname, ") Assert", fname, "Greater(min ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(&min, nil, 1)")
	g.P("}")
	g.P()

	// greater or equal
	g.P("func (x *", mname, ") Assert", fname, "GreaterOrEqual(min ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(&min, nil, 0)")
	g.P("}")
	g.P()

	// less
	g.P("func (x *", mname, ") Assert", fname, "Less(max ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(nil, &max, 1)")
	g.P("}")
	g.P()

	// less or equal
	g.P("func (x *", mname, ") Assert", fname, "LessOrEqual(max ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(nil, &max, 0)")
	g.P("}")
	g.P()
}

func generateString(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("// min, max: 0 for unlimited, rune count must be in [min, max].")
	g.P("func (x *", mname, ") Assert", fname, "RuneCountRange(min, max int) error {")
	g.P("v := x.Get", fname, "()")
	g.P("n := utf8.RuneCountInString(v)")

	g.P("if min >0 && n < min {")
	g.P(`return fmt.Errorf("`, fullName, ` rune count must be greater or equal than %d, count: %d, value: %q", min, n, v)`)
	g.P("}")

	g.P("if max >0 && n > max {")
	g.P(`return fmt.Errorf("`, fullName, ` rune count must be less or equal than %d, count: %d, value: %q", max, n, v)`)
	g.P("}")

	g.P("return nil")
	g.P("}")
	g.P()
}

func generateEnumerate(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	ename := g.QualifiedGoIdent(f.Enum.GoIdent)
	fullName := mname + "." + fname

	// slice
	g.P("func (x *", mname, ") Assert", fname, "InSlice(vs ...", ename, ") error {")
	g.P("v := x.Get", fname, "()")
	g.P("for _, valid := range vs {")
	g.P("if v == valid {")
	g.P("return nil")
	g.P("}")
	g.P("}")
	g.P(`return fmt.Errorf("`, fullName, ` must be in %v, value: %v", vs, v)`)
	g.P("}")
	g.P()

	// map
	g.P("func (x *", mname, ") Assert", fname, "InMap(m map[", ename, "]struct{}) error {")
	g.P("v := x.Get", fname, "()")
	g.P("if _, ok := m[v]; ok {")
	g.P("return nil")
	g.P("}")
	g.P("vs := make([]", ename, ", 0, len(m))")
	g.P("for valid := range m {")
	g.P("vs = append(vs, valid)")
	g.P("}")
	g.P(`return fmt.Errorf("`, fullName, ` must be in %v, value: %v", vs, v)`)
	g.P("}")
	g.P()
}
