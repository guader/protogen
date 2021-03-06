package plugins

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Validation struct{}

func (gen *Validation) GeneratedFilenameSuffix() string {
	return ".pb.enum.validation.go"
}

func (gen *Validation) Generate(g *protogen.GeneratedFile, enums []*protogen.Enum) error {
	for _, enum := range enums {
		name := enum.Desc.Name()
		g.P("// IsValid return whether the ", name, " is a defined value")
		g.P("func (x ", name, ") IsValid() bool {")
		g.P("_, ok := ", name, "_name[int32(x)]")
		g.P("return ok")
		g.P("}")
		g.P()
		g.P("// IsDefault return whether the ", name, " is the default value")
		g.P("func (x ", name, ") IsDefault() bool {")
		g.P("return x == 0")
		g.P("}")
		g.P()
	}
	return nil
}
