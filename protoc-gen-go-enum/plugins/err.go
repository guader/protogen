package plugins

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type Err struct{}

func (gen *Err) GeneratedFilenameSuffix() string {
	return ".pb.enum.err.go"
}

func (gen *Err) Generate(g *protogen.GeneratedFile, enums []*protogen.Enum) error {
	for _, enum := range enums {
		name := enum.Desc.Name()
		g.P("// Error implement error interface for ", name)
		g.P("func (x ", name, ") Error() string {")
		g.P("return x.String()")
		g.P("}")
		g.P()
	}
	return nil
}
