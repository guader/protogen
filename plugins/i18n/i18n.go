package i18n

import (
	"fmt"
	"sort"
	"strconv"
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

		generateForEnums(g, file.Enums)
		generateForMessages(g, file.Messages)
	}
	return nil
}

func getTranslations(opts *i18n.I18N) [][2]string {
	var list [][2]string
	opts.ProtoReflect().Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		lang := string(d.Name())
		text := v.String()
		list = append(list, [2]string{lang, text})
		return true
	})
	sort.SliceStable(list, func(i, j int) bool { return list[i][0] < list[j][0] })
	return list
}

func generateForEnums(g *protogen.GeneratedFile, es []*protogen.Enum) {
	for _, e := range es {
		enumOpts := pkg.ProtoGetExtension[i18n.EnumOptions](e.Desc.Options(), i18n.E_EnumOptions)
		if !enumOpts.GetEnable() {
			continue
		}

		enumName := e.GoIdent.GoName

		langSet := make(map[string]struct{})
		// Generate translation table.
		g.P("var ", enumName, "_i18n = map[", enumName, "]map[string]string{")
		for _, v := range e.Values {
			opts := pkg.ProtoGetExtension[i18n.I18N](v.Desc.Options(), i18n.E_Enum)
			if opts == nil {
				continue
			}

			ts := getTranslations(opts)
			if len(ts) == 0 {
				continue
			}

			g.P(g.QualifiedGoIdent(v.GoIdent), ": {")
			for _, t := range ts {
				lang := t[0]
				text := t[1]
				g.P(strconv.Quote(lang), ":", strconv.Quote(text), ",")
				langSet[lang] = struct{}{}
			}
			g.P("},")
		}
		g.P("}")
		g.P()

		// Generate method.
		g.P("func (x ", enumName, ") I18n(lang string) string {")
		g.P("if m, ok := ", enumName, "_i18n[x]; ok {")
		g.P("return m[lang]")
		g.P("}")
		g.P(`return ""`)
		g.P("}")
		g.P()
	}
}

func generateForMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
	for _, m := range ms {
		generateForEnums(g, m.Enums)
		generateForMessages(g, m.Messages)
	}
}

//func generateForEnums(g *protogen.GeneratedFile, es []*protogen.Enum) {
//	for _, e := range es {
//		enumOpts := pkg.ProtoGetExtension[i18n.EnumOptions](e.Desc.Options(), i18n.E_EnumOptions)
//		if !enumOpts.GetEnable() {
//			continue
//		}
//
//		enumName := e.GoIdent.GoName
//
//		langSet := make(map[string]struct{})
//		// Generate translation table.
//		g.P("var ", enumName, "_i18n = map[", enumName, "]map[string]string{")
//		for _, v := range e.Values {
//			opts := pkg.ProtoGetExtension[i18n.I18N](v.Desc.Options(), i18n.E_Enum)
//			if opts == nil {
//				continue
//			}
//
//			ts := getTranslations(opts)
//			if len(ts) == 0 {
//				continue
//			}
//
//			g.P(g.QualifiedGoIdent(v.GoIdent), ":{")
//			for _, t := range ts {
//				lang := t[0]
//				text := t[1]
//				g.P(strconv.Quote(lang), ":", strconv.Quote(text), ",")
//				langSet[lang] = struct{}{}
//			}
//			g.P("},")
//		}
//		g.P("}")
//		g.P()
//
//		if len(langSet) == 0 {
//			continue
//		}
//		var langs []string
//		for lang := range langSet {
//			langs = append(langs, lang)
//		}
//		sort.SliceStable(langs, func(i, j int) bool { return langs[i] < langs[j] })
//
//		// Generate method for enum.
//		for _, lang := range langs {
//			g.P("func (x ", enumName, ") I18n_", lang, "() string {")
//			g.P("if m, ok := ", enumName, "_i18n[x]; ok {")
//			g.P("return m[", strconv.Quote(lang), "]")
//			g.P("}")
//			g.P(`return ""`)
//			g.P("}")
//			g.P()
//		}
//
//		// Generate interface.
//		g.P("type I18n_", enumName, " interface {")
//		g.P("I18n_GetByValue(", enumName, ") string")
//		g.P("}")
//		g.P()
//
//		// Generate exported variables.
//		g.P("var (")
//		for _, lang := range langs {
//			g.P("I18n_", enumName, "_", lang, " I18n_", enumName, "= (*i18n_", enumName, "_", lang, ")(nil)")
//		}
//		g.P(")")
//
//		// Generate unexported implementations.
//		for _, lang := range langs {
//			g.P("type i18n_", enumName, "_", lang, " struct{}")
//			g.P()
//			g.P("func (*i18n_", enumName, "_", lang, ") I18n_GetByValue(v ", enumName, ") string {")
//			g.P("return v.I18n_", lang, "()")
//			g.P("}")
//			g.P()
//		}
//	}
//}

//func generateForMessages(g *protogen.GeneratedFile, ms []*protogen.Message) {
//	for _, m := range ms {
//		generateForEnums(g, m.Enums)
//		generateForMessages(g, m.Messages)
//
//		messageOpts := pkg.ProtoGetExtension[i18n.MessageOptions](m.Desc.Options(), i18n.E_MessageOptions)
//		if !messageOpts.GetEnable() {
//			continue
//		}
//
//		messageName := m.GoIdent.GoName
//
//		// Gather field information.
//		langSet := make(map[string]struct{})
//		transesByLang := make(map[string]map[string]string) // map[lang]map[field]text
//		for _, f := range m.Fields {
//			opts := pkg.ProtoGetExtension[i18n.I18N](f.Desc.Options(), i18n.E_Field)
//			for _, t := range getTranslations(opts) {
//				lang := t[0]
//				text := t[1]
//				langSet[lang] = struct{}{}
//				if _, ok := transesByLang[lang]; !ok {
//					transesByLang[lang] = make(map[string]string)
//				}
//				transesByLang[lang][f.GoName] = text
//			}
//		}
//		if len(langSet) == 0 {
//			continue
//		}
//
//		// Generate interface.
//		g.P("type I18n_", messageName, " interface {")
//		for _, f := range m.Fields {
//			g.P(f.GoName, "() string")
//
//		}
//		g.P("}")
//		g.P()
//
//		var langs []string
//		for lang := range langSet {
//			langs = append(langs, lang)
//		}
//		sort.SliceStable(langs, func(i, j int) bool { return langs[i] < langs[j] })
//
//		// Generate exported variables.
//		g.P("var (")
//		for _, lang := range langs {
//			g.P("I18n_", messageName, "_", lang, " I18n_", messageName, " = (*i18n_", messageName, "_", lang, ")(nil)")
//		}
//		g.P(")")
//		g.P()
//
//		// Generate implementations.
//		for _, lang := range langs {
//			g.P("type i18n_", messageName, "_", lang, " struct{}")
//			g.P()
//			for _, f := range m.Fields {
//				fieldName := f.GoName
//				var text string
//				if textByFieldName, ok := transesByLang[lang]; ok {
//					text = textByFieldName[fieldName]
//				}
//				g.P("func (*i18n_", messageName, "_", lang, ")", fieldName, "() string { return ", strconv.Quote(text), "}")
//				g.P()
//			}
//		}
//	}
//}
