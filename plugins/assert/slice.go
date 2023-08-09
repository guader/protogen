package assert

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateSlice(g *protogen.GeneratedFile, m *protogen.Message, f *protogen.Field) {
	mname := m.GoIdent.GoName
	fname := f.GoName
	fullName := mname + "." + fname

	g.P("// min, max: 0 for unlimited, item count must be in range [min, max].")
	g.P("func (x *", mname, ") Assert", fname, "_ItemCountRange(min, max int) AssertFunc {")
	g.P("return func() error {")
	g.P(fmt.Sprintf("return AssertSliceLengthRange(%q, x.Get%s(), min, max)", fullName, fname))
	g.P("}")
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
