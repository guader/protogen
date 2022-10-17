package main

import (
	"flag"
	"strings"

	gengo "google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"

	"github.com/guader/protogen/protoc-gen-go-svc/plugins"
	"github.com/guader/protogen/version"
)

const (
	pluginMethodEnum = "methodEnum"
)

var (
	flags        flag.FlagSet
	flagPlugins  string
	flagServices string
)

type generator interface {
	Generate(*protogen.GeneratedFile, []*protogen.Service) error
	GeneratedFilenameSuffix() string
}

func main() {
	flags.StringVar(&flagPlugins, "plugins", "", "plugins for code generation, example: methodEnum, available: methodEnum.")
	flags.StringVar(&flagServices, "services", "", "generate code for services of specified names, for all if not specified. example: A+B, generate code for service A and B.")
	protogen.Options{ParamFunc: flags.Set}.Run(generate)
}

func generate(plugin *protogen.Plugin) error {
	plugin.SupportedFeatures = gengo.SupportedFeatures

	var specifiedGenerators []generator
	for _, s := range strings.Split(flagPlugins, "+") {
		switch s {
		case pluginMethodEnum:
			specifiedGenerators = append(specifiedGenerators, &plugins.MethodEnum{})
		}
	}
	if len(specifiedGenerators) == 0 {
		return nil
	}

	var specifiedServices map[string]struct{}
	if flagServices != "" {
		specifiedServices = make(map[string]struct{})
		for _, s := range strings.Split(flagServices, "+") {
			specifiedServices[s] = struct{}{}
		}
	}

	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		var services []*protogen.Service
		if flagServices != "" {
			for _, service := range file.Services {
				_, ok := specifiedServices[string(service.Desc.Name())]
				if ok {
					services = append(services, service)
				}
			}
		} else {
			services = file.Services
		}
		if len(services) == 0 {
			continue
		}

		for _, gen := range specifiedGenerators {
			g := plugin.NewGeneratedFile(file.GeneratedFilenamePrefix+gen.GeneratedFilenameSuffix(), file.GoImportPath)
			g.P("// Code generated by protoc-gen-go-svc. DO NOT EDIT.")
			g.P("// version: ", version.Version)
			g.P("// source: ", file.Desc.Path())
			g.P()
			g.P("package ", file.GoPackageName)
			g.P()
			if err := gen.Generate(g, services); err != nil {
				return err
			}
		}
	}
	return nil
}
