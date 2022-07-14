package plugins

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Validation struct{}

func (gen *Validation) Version() string {
	return "v0.0.1"
}

func (gen *Validation) GeneratedFilenameSuffix() string {
	return ".validation.go"
}

func (gen *Validation) Generate(g *protogen.GeneratedFile, enums []*protogen.Enum) error {
	for _, enum := range enums {
		name := enum.Desc.Name()
		g.P("// IsValid returns whether the ", name, " is a defined value")
		g.P("func (x ", name, ") IsValid() bool {")
		g.P("_, ok := ", name, "_name[int32(x)]")
		g.P("return ok")
		g.P("}")
		g.P()
		g.P("// IsDefault returns whether the ", name, " is the default value")
		g.P("func (x ", name, ") IsDefault() bool {")
		g.P("return x == 0")
		g.P("}")
		g.P()
	}
	return nil
}
