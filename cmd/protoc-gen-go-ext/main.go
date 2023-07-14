package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"github.com/guader/protogen/plugins/consts"
	"github.com/guader/protogen/plugins/enums"
	"github.com/guader/protogen/plugins/errorer"
	"github.com/guader/protogen/plugins/i18n"
	"github.com/guader/protogen/plugins/setter"
	"github.com/guader/protogen/plugins/validator"
)

func main() {
	protogen.Options{}.Run(generate)
}

func generate(plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	for _, f := range []func(*protogen.Plugin) error{
		setter.Generate,
		validator.Generate,
		errorer.Generate,
		consts.Generate,
		i18n.Generate,
		enums.Generate,
	} {
		if err := f(plugin); err != nil {
			return err
		}
	}
	return nil
}
