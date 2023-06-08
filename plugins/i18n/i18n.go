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

type language = protoreflect.Name

func generateEnums(g *protogen.GeneratedFile, es []*protogen.Enum) {
	for _, e := range es {
		transesByLang := make(map[language]map[string]string)
		for _, v := range e.Values {
			opts := pkg.ProtoGetExtension[i18n.I18N](v.Desc.Options(), i18n.E_Enum)
			if opts == nil {
				continue
			}
			value := g.QualifiedGoIdent(v.GoIdent)
			opts.ProtoReflect().Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				lang := d.Name()
				transes, ok := transesByLang[lang]
				if !ok {
					transes = make(map[string]string)
					transesByLang[lang] = transes
				}
				transes[value] = v.String()
				return true
			})
		}
		if len(transesByLang) == 0 {
			continue
		}

		typ := g.QualifiedGoIdent(e.GoIdent)
		Prefix := "I18n_" + typ // Export.
		prefix := "i18n_" + typ // Unexport.
		g.P("type ", Prefix, " interface {")
		g.P("I18n_GetByValue(v ", typ, ") string")
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
			g.P("func (x ", typ, ") I18n_", lang, "() string {")
			g.P("return map[", typ, "]string{")
			transes := transesByLang[lang]
			for _, v := range e.Values {
				value := g.QualifiedGoIdent(v.GoIdent)
				text, ok := transes[value]
				if !ok {
					continue
				}
				g.P(fmt.Sprintf(`%s: %q,`, value, text))
			}
			g.P("}[x]")
			g.P("}")

			g.P()
			g.P(fmt.Sprintf(`var %s_%s %s = (*%s_%s)(nil)`, Prefix, lang, Prefix, prefix, lang))
			g.P()
			g.P(fmt.Sprintf(`type %s_%s struct{}`, prefix, lang))
			g.P()
			g.P(fmt.Sprintf(`func (*%s_%s) I18n_GetByValue(v %s) string {`,
				prefix, lang, typ))
			g.P("return v.I18n_", lang, "()")
			//g.P("switch v {")
			//for _, v := range e.Values {
			//	value := g.QualifiedGoIdent(v.GoIdent)
			//	text, ok := transes[value]
			//	if !ok {
			//		continue
			//	}
			//	g.P("case ", value, ":")
			//	g.P(fmt.Sprintf(`return %q`, text))
			//}
			//g.P("default:")
			//g.P(`return ""`)
			//g.P("}")
			g.P("}")
		}
	}
}

func generateMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		generateEnums(g, m.Enums)
		generateMessages(g, m.Messages)

		fields := make([]string, 0)
		transesByLang := make(map[language]map[string]string)
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
					transes = make(map[string]string)
					transesByLang[lang] = transes
				}
				transes[field] = v.String()
				return true
			})
		}
		if len(fields) == 0 {
			continue
		}

		typ := g.QualifiedGoIdent(m.GoIdent)
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
}
