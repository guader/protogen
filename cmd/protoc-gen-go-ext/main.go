package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/config"
	"github.com/guader/protogen/plugins/setter"
)

func main() {
	conf := &config.Config{}
	flags := flag.FlagSet{}
	flags.BoolVar(&conf.Setter.Enabled, "setter", false, "")
	flags.StringVar(&conf.Setter.Suffix, "setter_suffix", "setter", "")

	protogen.Options{ParamFunc: flags.Set}.Run(generate(conf))
}

type pluginFunc = func(*config.Config, *protogen.Plugin) error

func generate(conf *config.Config) func(*protogen.Plugin) error {
	return func(plugin *protogen.Plugin) error {
		for _, f := range []pluginFunc{
			setter.Generate,
		} {
			if err := f(conf, plugin); err != nil {
				return err
			}
		}
		return nil
	}
}
