package assert

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pkg"
)

func generateNumber(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	typ, _ := pkg.FieldGoType(g, f)
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	// range
	g.P("// open: 0 for both close, 1 for left open only, 2 for right open only, 3 for both open.")
	g.P("func (x *", mname, ") Assert_", fname, "_Range(min, max *", typ, ", open byte) AssertFunc {")
	g.P("return func() error {")
	g.P(fmt.Sprintf("return AssertNumberRange(%q, x.Get%s(), min, max, open)", fullName, fname))
	g.P("}")
	g.P("}")
	g.P()

	// greater
	g.P("func (x *", mname, ") Assert_", fname, "_GT(min ", typ, ") AssertFunc { return x.Assert_", fname,
		"_Range(&min, nil, 1) }")
	// greater or equal
	g.P("func (x *", mname, ") Assert_", fname, "_GE(min ", typ, ") AssertFunc { return x.Assert_", fname,
		"_Range(&min, nil, 0) }")
	// less
	g.P("func (x *", mname, ") Assert_", fname, "_LT(max ", typ, ") AssertFunc { return x.Assert_", fname,
		"_Range(nil, &max, 1) }")
	// less or equal
	g.P("func (x *", mname, ") Assert_", fname, "_LE(max ", typ, ") AssertFunc { return x.Assert_", fname,
		"_Range(nil, &max, 0) }")
	g.P()
}
