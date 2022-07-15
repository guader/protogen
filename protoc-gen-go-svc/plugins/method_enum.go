package plugins

import (
	"fmt"
	"strconv"

	"google.golang.org/protobuf/compiler/protogen"
)

type MethodEnum struct{}

func (gen *MethodEnum) GeneratedFilenameSuffix() string {
	return ".pb.svc.method_enum.go"
}

func (gen *MethodEnum) Generate(g *protogen.GeneratedFile, services []*protogen.Service) error {
	for _, service := range services {
		serviceName := service.Desc.Name()
		serviceFullName := service.Desc.FullName()
		g.P("const (")
		for _, method := range service.Methods {
			methodName := method.Desc.Name()
			k := fmt.Sprintf("%s_Method_%s", serviceName, methodName)
			v := strconv.Quote(fmt.Sprintf("/%s/%s", serviceFullName, methodName))
			g.P(fmt.Sprintf("%s = %s", k, v))
		}
		g.P(")")
		g.P()
	}
	return nil
}
