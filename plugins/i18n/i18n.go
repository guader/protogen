package i18n

import (
	"fmt"
	"sort"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/guader/protogen/pb/i18n"
	"github.com/guader/protogen/pkg"
	"github.com/guader/protogen/version"
)

func Generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		fileOpts := pkg.ProtoGetExtension[i18n.FileOptions](file.Desc.Options(), i18n.E_FileOptions)
		if !fileOpts.GetEnable() {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.i18n"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.go", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "i18n", file.Desc.Path()))
		g.P("package ", file.GoPackageName)

		generateEnums(g, file.Enums)
		generateMessages(g, file.Messages)
	}
	return nil
}

type (
	language     = protoreflect.Name
	translations = map[string]string
)

func generateEnums(g *protogen.GeneratedFile, es []*protogen.Enum) {
	for _, e := range es {
		fields := make([]string, 0)
		transesByLang := make(map[language]translations)
		for _, v := range e.Values {
			opts := pkg.ProtoGetExtension[i18n.I18N](v.Desc.Options(), i18n.E_Enum)
			if opts == nil {
				continue
			}
			field := exportName(string(v.Desc.Name())) // Want shorter name but exported.
			fields = append(fields, field)
			opts.ProtoReflect().Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				lang := d.Name()
				transes, ok := transesByLang[lang]
				if !ok {
					transes = make(translations)
					transesByLang[lang] = transes
				}
				transes[field] = v.String()
				return true
			})
		}
		if len(fields) == 0 {
			continue
		}
		generateI18n(g, e.GoIdent.GoName, fields, transesByLang)
	}
}

func generateMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		fields := make([]string, 0)
		transesByLang := make(map[language]translations)
		for _, f := range m.Fields {
			opts := pkg.ProtoGetExtension[i18n.I18N](f.Desc.Options(), i18n.E_Field)
			if opts == nil {
				continue
			}
			field := f.GoName
			fields = append(fields, field)
			opts.ProtoReflect().Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				lang := d.Name()
				transes, ok := transesByLang[lang]
				if !ok {
					transes = make(translations)
					transesByLang[lang] = transes
				}
				transes[field] = v.String()
				return true
			})
		}
		if len(fields) == 0 {
			continue
		}
		generateI18n(g, m.GoIdent.GoName, fields, transesByLang)
		generateEnums(g, m.Enums)
		generateMessages(g, m.Messages)
	}
}

func generateI18n(g *protogen.GeneratedFile, typ string, fields []string, transesByLang map[language]translations) {
	Prefix := "I18n_" + typ // Export.
	prefix := "i18n_" + typ // Unexport.
	g.P()
	g.P("type ", Prefix, " interface {")
	for _, field := range fields {
		g.P(field, "() string")
	}
	g.P("}")

	langs := make([]language, 0, len(transesByLang))
	for lang := range transesByLang {
		langs = append(langs, lang)
	}
	sort.SliceStable(langs, func(i, j int) bool {
		return langs[i] < langs[j]
	})
	for _, lang := range langs {
		g.P()
		g.P(fmt.Sprintf(`var %s_%s %s = (*%s_%s)(nil)`, Prefix, lang, Prefix, prefix, lang))
		g.P()
		g.P(fmt.Sprintf(`type %s_%s struct{}`, prefix, lang))
		g.P()
		transes := transesByLang[lang]
		for _, field := range fields {
			g.P(fmt.Sprintf(`func (*%s_%s) %s() string { return %q }`, prefix, lang, field, transes[field]))
		}
	}
}

func exportName(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}
