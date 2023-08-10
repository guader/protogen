package assert

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateBytes(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("// min, max: 0 for unlimited, bytes count must be in range [min, max].")
	g.P("func (x *", mname, ") Assert_", fname, "_BytesCountRange(min, max int) AssertFunc {")
	g.P("return func() error {")
	g.P(fmt.Sprintf("return AssertSliceLengthRange(%q, x.Get%s(), min, max)", fullName, fname))
	g.P("}")
	g.P("}")
	g.P()
}
