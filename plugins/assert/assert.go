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
			cg.P("import (")
			cg.P(`"fmt"`)
			cg.P(`"unicode/utf8"`)
			cg.P(")")
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

const (
	assertFuncTypeName = "AssertFunc"

	assertNumberRangeFuncName      = "AssertNumberRange"
	assertNumberInSliceFuncName    = "AssertNumberInSlice"
	assertNumberInMapFuncName      = "AssertNumberInMap"
	assertRuneCountRangeFuncName   = "AssertRuneCountRange"
	assertSliceLengthRangeFuncName = "AssertSliceLengthRange"
)

func generateCommon(g *protogen.GeneratedFile) {
	// interface
	g.P(`type NumberType interface {
	~int32 | ~int64 | ~uint32 | ~uint64 | ~float32 | ~float64
}`)
	g.P()

	// types
	g.P("type ", assertFuncTypeName, " = func() error")
	g.P()

	// number range
	g.P("// open: 0 for both close, 1 for left open only, 2 for right open only, 3 for both open.")
	g.P("func ", assertNumberRangeFuncName, "[T NumberType](name string, v T, min, max *T, open byte) error {")
	g.P(`if min != nil {
		if open&1 == 0 {
			// left close
			if v < *min {
				return fmt.Errorf("%s must be greater or equal than %v, value: %v", name, *min, v)
			}
		} else {
			// left open
			if v <= *min {
				return fmt.Errorf("%s must be greater than %v, value: %v", name, *min, v)
			}
		}
	}
	if max != nil {
		if open&2 == 0 {
			// right close
			if v > *max {
				return fmt.Errorf("%s must be less or equal than %v, value: %v", name, *max, v)
			}
		} else {
			// right open
			if v >= *max {
				return fmt.Errorf("%s must be less than %v, value: %v", name, *max, v)
			}
		}
	}
	return nil`)
	g.P("}")
	g.P()

	// number in slice
	g.P("func ", assertNumberInSliceFuncName, "[T NumberType](name string, v T, vs ...T) error {")
	g.P(`for _, valid := range vs {
		if valid == v {
			return nil
		}
	}
	return fmt.Errorf("%s must be in %v, value: %v", name, vs, v)`)
	g.P("}")
	g.P()

	// number in map
	g.P("func ", assertNumberInMapFuncName, "[T NumberType](name string, v T, m map[T]struct{}) error {")
	g.P(`if _, ok := m[v]; ok {
		return nil
	}
	vs := make([]T, 0, len(m))
	for valid := range m {
		vs = append(vs, valid)
	}
	return fmt.Errorf("%s must be in %v, value: %v", name, vs, v)`)
	g.P("}")
	g.P()

	// string rune count
	g.P("// min, max: 0 for unlimited, rune count must be in range [min, max].")
	g.P("func ", assertRuneCountRangeFuncName, "(name string, v string, min, max int) error {")
	g.P(`n := utf8.RuneCountInString(v)
		if min > 0 && n < min {
			return fmt.Errorf("%s rune count must be greater or equal than %d, count: %d, value: %q", name, min, n, v)
		}
		if max > 0 && n > max {
			return fmt.Errorf("%s rune count must be less or equal than %d, count: %d, value: %q", name, max, n, v)
		}
		return nil`)
	g.P("}")
	g.P()

	// slice length
	g.P("// min, max: 0 for unlimited, slice length must be in range [min, max].")
	g.P("func ", assertSliceLengthRangeFuncName, "[T any](name string, v []T, min, max int) error {")
	g.P(`n := len(v)
	if min > 0 && n < min {
		return fmt.Errorf("%s slice length must be greater or equal than %d, count: %d", name, min, n)
	}
	if max > 0 && n > max {
		return fmt.Errorf("%s slice length must be less or equal than %d, count: %d", name, max, n)
	}
	return nil`)
	g.P("}")
	g.P()
}

func generateNumber(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	typ, _ := pkg.FieldGoType(g, f)
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	// range
	g.P("// open: 0 for both close, 1 for left open only, 2 for right open only, 3 for both open.")
	g.P("func (x *", mname, ") Assert", fname, "Range(min, max *", typ, ", open byte) error {")
	g.P(fmt.Sprintf("return %s(%q, x.Get%s(), min, max, open)", assertNumberRangeFuncName, fullName, fname))
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "RangeFunc(min, max *", typ, ", open byte) ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "Range(min, max, open) }")
	g.P("}")
	g.P()

	// greater
	g.P("func (x *", mname, ") Assert", fname, "Greater(min ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(&min, nil, 1)")
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "GreaterFunc(min ", typ, ") ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "Greater(min) }")
	g.P("}")
	g.P()

	// greater or equal
	g.P("func (x *", mname, ") Assert", fname, "GreaterOrEqual(min ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(&min, nil, 0)")
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "GreaterOrEqualFunc(min ", typ, ") ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "GreaterOrEqual(min) }")
	g.P("}")
	g.P()

	// less
	g.P("func (x *", mname, ") Assert", fname, "Less(max ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(nil, &max, 1)")
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "LessFunc(max ", typ, ") ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "Less(max) }")
	g.P("}")
	g.P()

	// less or equal
	g.P("func (x *", mname, ") Assert", fname, "LessOrEqual(max ", typ, ") error {")
	g.P("return x.Assert", fname, "Range(nil, &max, 0)")
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "LessOrEqualFunc(max ", typ, ") ", assertFuncTypeName, " {")
	g.P("return func() error { return  x.Assert", fname, "LessOrEqual(max) }")
	g.P("}")
	g.P()
}

func generateString(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("// min, max: 0 for unlimited, rune count must be in range [min, max].")
	g.P("func (x *", mname, ") Assert", fname, "RuneCountRange(min, max int) error {")
	g.P(fmt.Sprintf("return %s(%q, x.Get%s(), min, max)", assertRuneCountRangeFuncName, fullName, fname))
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "RuneCountRangeFunc(min, max int) ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "RuneCountRange(min, max) }")
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
	g.P(fmt.Sprintf("return %s(%q, x.Get%s(), vs...)", assertNumberInSliceFuncName, fullName, fname))
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "InSliceFunc(vs ...", ename, ") ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "InSlice(vs...) }")
	g.P("}")
	g.P()

	// map
	g.P("func (x *", mname, ") Assert", fname, "InMap(m map[", ename, "]struct{}) error {")
	g.P(fmt.Sprintf("return %s(%q, x.Get%s(), m)", assertNumberInMapFuncName, fullName, fname))
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "InMapFunc(m map[", ename, "]struct{}) ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "InMap(m) }")
	g.P("}")
	g.P()
}

func generateBytes(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("// min, max: 0 for unlimited, bytes count must be in range [min, max].")
	g.P("func (x *", mname, ") Assert", fname, "BytesCountRange(min, max int) error {")
	g.P(fmt.Sprintf("return %s(%q, x.Get%s(), min, max)", assertSliceLengthRangeFuncName, fullName, fname))
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "BytesCountRangeFunc(min, max int) ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "BytesCountRange(min, max) }")
	g.P("}")
	g.P()
}

func generateSlice(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("// min, max: 0 for unlimited, item count must be in range [min, max].")
	g.P("func (x *", mname, ") Assert", fname, "ItemCountRange(min, max int) error {")
	g.P(fmt.Sprintf("return %s(%q, x.Get%s(), min, max)", assertSliceLengthRangeFuncName, fullName, fname))
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "ItemCountRangeFunc(min, max int) ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "ItemCountRange(min, max) }")
	g.P("}")
	g.P()

	// nillable
	//if f.Desc.Kind() == protoreflect.BytesKind ||
	//	f.Desc.Kind() == protoreflect.MessageKind ||
	//	f.Desc.Kind() == protoreflect.GroupKind {
	//	g.P("func (x *", mname, ") Assert", fname, "ItemNonNil() error {")
	//	g.P("for i, item := range x.Get", fname, "() {")
	//	g.P("if item == nil {")
	//	g.P(`return fmt.Errorf("`, fullName, ` item must not be nil, index: %d", i)`)
	//	g.P("}")
	//	g.P("}")
	//	g.P("return nil")
	//	g.P("}")
	//	g.P()
	//}
}

func generateNillable(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("func (x *", mname, ") Assert", fname, "NonNil() error {")
	g.P("if x == nil || x.", fname, " == nil {")
	g.P(`return fmt.Errorf("`, fullName, ` must not be nil")`)
	g.P("}")
	g.P("return nil")
	g.P("}")
	g.P()
	g.P("func (x *", mname, ") Assert", fname, "NonNilFunc() ", assertFuncTypeName, " {")
	g.P("return func() error { return x.Assert", fname, "NonNil() }")
	g.P("}")
	g.P()
}
