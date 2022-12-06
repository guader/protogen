package plugins

import (
	"fmt"
	"sort"
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
		urlSet := make([]string, 0)
		methodByUrl := make(map[string]string)
		methodSet := make([]string, 0)
		urlByMethod := make(map[string]string)
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

				mk := fmt.Sprintf("%s_Method_%s", serviceName, methodName)
				urlSet = append(urlSet, k)
				methodByUrl[k] = mk
				methodSet = append(methodSet, mk)
				urlByMethod[mk] = k
			}
		}
		g.P(")")
		g.P()

		g.P("var ", serviceName, "_MethodByURL = map[string]string{")
		sort.Strings(urlSet)
		for _, u := range urlSet {
			g.P(u, ": ", methodByUrl[u], ",")
		}
		g.P("}")
		g.P()

		g.P("var ", serviceName, "_URLByMethod = map[string]string{")
		sort.Strings(methodSet)
		for _, m := range methodSet {
			g.P(m, ": ", urlByMethod[m], ",")
		}
		g.P("}")
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
