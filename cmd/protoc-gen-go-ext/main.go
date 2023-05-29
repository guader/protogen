package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/plugins/errorer"
	"github.com/guader/protogen/plugins/setter"
	"github.com/guader/protogen/plugins/validator"
)

func main() {
	protogen.Options{}.Run(generate)
}

func generate(plugin *protogen.Plugin) error {
	for _, f := range []func(*protogen.Plugin) error{
		setter.Generate,
		validator.Generate,
		errorer.Generate,
	} {
		if err := f(plugin); err != nil {
			return err
		}
	}
	return nil
}
