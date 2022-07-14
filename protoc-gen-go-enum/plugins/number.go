package plugins

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Number struct{}

func (gen *Number) GeneratedFilenameSuffix() string {
	return ".pb.enum.number.go"
}

func (gen *Number) Generate(g *protogen.GeneratedFile, enums []*protogen.Enum) error {
	for _, enum := range enums {
		name := enum.Desc.Name()
		g.P("// Int32 convert ", name, " to int32")
		g.P("func (x ", name, ") Int32() int32 {")
		g.P("return int32(x)")
		g.P("}")
		g.P()
		g.P("// Int64 convert ", name, " to int64")
		g.P("func (x ", name, ") Int64() int64 {")
		g.P("return int64(x)")
		g.P("}")
		g.P()
	}
	return nil
}
