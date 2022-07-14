package plugins

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type GrpcStatus struct{}

func (gen *GrpcStatus) Version() string {
	return "v0.0.1"
}

func (gen *GrpcStatus) GeneratedFilenameSuffix() string {
	return ".pb.grpc_status.go"
}

func (gen *GrpcStatus) Generate(g *protogen.GeneratedFile, enums []*protogen.Enum) error {
	g.P("import (")
	g.P(`"google.golang.org/grpc/codes"`)
	g.P(`"google.golang.org/grpc/status"`)
	g.P(")")
	g.P()
	for _, enum := range enums {
		name := enum.Desc.Name()
		g.P("// GrpcErr apply status.Error() to ", name, " with ", name, ".String() as message")
		g.P("func (x ", name, ") GrpcErr() error {")
		g.P("return status.Error(codes.Code(x), x.String())")
		g.P("}")
		g.P()
		g.P("// GrpcError apply status.Error() to ", name)
		g.P("func (x ", name, ") GrpcError(msg string) error {")
		g.P("return status.Error(codes.Code(x), msg)")
		g.P("}")
		g.P()
		g.P("// GrpcErrorf apply status.Errorf() to ", name)
		g.P("func (x ", name, ") GrpcErrorf(format string, a ...interface{}) error {")
		g.P("return status.Errorf(codes.Code(x),format, a...)")
		g.P("}")
		g.P()
	}
	return nil
}
