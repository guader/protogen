package assert

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateEnumerate(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	ename := g.QualifiedGoIdent(f.Enum.GoIdent)
	fullName := mname + "." + fname

	// slice
	g.P("func (x *", mname, ") Assert", fname, "_InSlice(vs ...", ename, ") AssertFunc {")
	g.P("return func() error {")
	g.P(fmt.Sprintf("return AssertNumberInSlice(%q, x.Get%s(), vs...)", fullName, fname))
	g.P("}")
	g.P("}")
	g.P()

	// map
	g.P("func (x *", mname, ") Assert", fname, "_InMap(m map[", ename, "]struct{}) AssertFunc {")
	g.P("return func() error {")
	g.P(fmt.Sprintf("return AssertNumberInMap(%q, x.Get%s(), m)", fullName, fname))
	g.P("}")
	g.P("}")
	g.P()
}
