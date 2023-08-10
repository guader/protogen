package assert

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func generateNillable(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("func (x *", mname, ") Assert_", fname, "_NotNil() AssertFunc {")
	g.P("return func() error {")
	g.P("if x == nil || x.", fname, " == nil {")
	g.P(`return fmt.Errorf("`, fullName, ` must not be nil")`)
	g.P("}")
	g.P("return nil")
	g.P("}")
	g.P("}")
	g.P()
}
