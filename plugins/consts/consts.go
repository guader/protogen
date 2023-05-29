package consts

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/guader/protogen/pb/consts"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

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
		g.P(pkg.RenderPackageComments(version.Version, "consts", file.Desc.Path(), string(file.GoPackageName)))

		for _, service := range file.Services {

			var (
				urlKV    []string
				methodKV []string
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
				urlKV = append(urlKV, fmt.Sprintf(`%s = %s`, urlKey, strconv.Quote(path)))

				methodKey := fmt.Sprintf(`%s_%s_FullMethodName`, service.GoName, method.Desc.Name())
				methodKV = append(methodKV, fmt.Sprintf(`%s: %s,`, methodKey, urlKey))
			}

			g.P("// URLs.")
			g.P("const (")
			for _, kv := range urlKV {
				g.P(kv)
			}
			g.P(")")

			g.P("// URLByMethod.")
			g.P(fmt.Sprintf(`var %s_URLByMethod = map[string]string{`, service.GoName))
			for _, kv := range methodKV {
				g.P(kv)
			}
			g.P("}")
		}
	}
	return nil
}
