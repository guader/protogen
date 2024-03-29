package consts

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/pb/consts"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[consts.FileOptions](file.Desc.Options(), consts.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.consts"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "consts", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()

		for _, service := range file.Services {
			var (
				urlKV         []string
				urlByMethodKV []string
				methodByURLKV []string
			)
			for _, method := range service.Methods {
				opts := pkg.ProtoGetExtension[annotations.HttpRule](method.Desc.Options(), annotations.E_Http)
				if opts == nil || opts.Pattern == nil {
					continue
				}

				var (
					methodType string
					path       string
				)
				switch opts.Pattern.(type) {
				case *annotations.HttpRule_Get:
					methodType = "GET"
					path = opts.GetGet()
				case *annotations.HttpRule_Put:
					methodType = "PUT"
					path = opts.GetPut()
				case *annotations.HttpRule_Post:
					methodType = "POST"
					path = opts.GetPost()
				case *annotations.HttpRule_Delete:
					methodType = "DELETE"
					path = opts.GetDelete()
				case *annotations.HttpRule_Patch:
					methodType = "PATCH"
					path = opts.GetPatch()
				default:
					continue
				}

				urlKey := fmt.Sprintf(`%s_URL_%s_%s`, service.GoName, methodType, method.Desc.Name())
				methodKey := fmt.Sprintf(`%s_%s_FullMethodName`, service.GoName, method.Desc.Name())

				urlKV = append(urlKV, fmt.Sprintf(`%s = %s`, urlKey, strconv.Quote(path)))
				urlByMethodKV = append(urlByMethodKV, fmt.Sprintf(`%s: %s,`, methodKey, urlKey))
				methodByURLKV = append(methodByURLKV, fmt.Sprintf(`%s: %s,`, urlKey, methodKey))
			}

			g.P("// URLs.")
			g.P("const (")
			for _, kv := range urlKV {
				g.P(kv)
			}
			g.P(")")

			g.P("// URLByMethod.")
			g.P(fmt.Sprintf(`var %s_URLByMethod = map[string]string{`, service.GoName))
			for _, kv := range urlByMethodKV {
				g.P(kv)
			}
			g.P("}")

			g.P("// MethodByURL.")
			g.P(fmt.Sprintf(`var %s_MethodByURL = map[string]string{`, service.GoName))
			for _, kv := range methodByURLKV {
				g.P(kv)
			}
			g.P("}")
		}
	}
	return nil
}
