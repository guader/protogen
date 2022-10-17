package plugins

import (
	"fmt"
	"strconv"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

type MethodEnum struct{}

func (gen *MethodEnum) GeneratedFilenameSuffix() string {
	return ".pb.svc.method_enum.go"
}

func (gen *MethodEnum) Generate(g *protogen.GeneratedFile, services []*protogen.Service) error {
	for _, service := range services {
		serviceName := uppercase(string(service.Desc.Name()))
		serviceFullName := service.Desc.FullName()
		g.P("// methods")
		g.P("const (")
		for _, method := range service.Methods {
			methodName := method.Desc.Name()
			k := fmt.Sprintf("%s_Method_%s", serviceName, methodName)
			v := strconv.Quote(fmt.Sprintf("/%s/%s", serviceFullName, methodName))
			g.P(fmt.Sprintf("%s = %s", k, v))
		}
		g.P(")")
		g.P()
		g.P("// urls")
		g.P("const (")
		for _, method := range service.Methods {
			opts := method.Desc.Options()
			if opts == nil {
				continue
			}
			httpRule, ok := proto.GetExtension(opts, annotations.E_Http).(*annotations.HttpRule)
			if !ok {
				continue
			}
			httpMethods := map[string]string{
				"GET":    httpRule.GetGet(),
				"POST":   httpRule.GetPost(),
				"PUT":    httpRule.GetPut(),
				"PATCH":  httpRule.GetPatch(),
				"DELETE": httpRule.GetDelete(),
			}
			methodName := method.Desc.Name()
			for m, u := range httpMethods {
				if u == "" {
					continue
				}
				k := fmt.Sprintf("%s_URL_%s_%s", serviceName, m, methodName)
				v := strconv.Quote(u)
				g.P(fmt.Sprintf("%s = %s", k, v))
			}
		}
		g.P(")")
		g.P()
	}
	return nil
}

func uppercase(s string) string {
	if len(s) > 0 && s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s
}
