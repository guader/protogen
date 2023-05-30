package errorer

import (
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/guader/protogen/pb/errorer"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[errorer.FileOptions](file.Desc.Options(), errorer.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.errorer"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "errorer", file.Desc.Path()))
		g.P("package ", file.GoPackageName)
		g.P()

		for _, enum := range file.Enums {
			enumOpts := pkg.ProtoGetExtension[errorer.EnumOptions](enum.Desc.Options(), errorer.E_EnumOptions)
			if !enumOpts.GetEnable() {
				continue
			}

			g.P("func (x ", enum.Desc.Name(), ") Error() string {")
			g.P("var m = map[int32]string{")
			for _, value := range enum.Values {
				errOpts := pkg.ProtoGetExtension[errorer.Error](value.Desc.Options(), errorer.E_Err)
				if errOpts == nil || errOpts.Message == nil {
					continue
				}
				g.P(fmt.Sprintf(`%d: %s,`, value.Desc.Number(), strconv.Quote(errOpts.GetMessage())))
			}

			g.P("}")
			g.P("if msg, ok := m[int32(x)]; ok {")
			g.P("return msg")
			g.P("}")
			g.P("return x.String()")
			g.P("}")
		}
	}
	return nil
}
