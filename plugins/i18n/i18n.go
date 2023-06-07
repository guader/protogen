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

		objectsByLang := make(map[string]map[string][][2]string)
		translationsForEnums(objectsByLang, file.Enums)
		translationsForMessages(objectsByLang, file.Messages)
		translations := make([]translation, 0, len(objectsByLang))
		for lang, fieldsByTypeName := range objectsByLang {
			objects := make([]object, 0, len(fieldsByTypeName))
			for typeName, fields := range fieldsByTypeName {
				objects = append(objects, object{
					name:   typeName,
					fields: fields,
				})
			}
			sort.SliceStable(objects, func(i, j int) bool {
				return objects[i].name < objects[j].name
			})
			translations = append(translations, translation{
				lang:    lang,
				objects: objects,
			})
		}
		sort.SliceStable(translations, func(i, j int) bool {
			return translations[i].lang < translations[j].lang
		})
		if len(translations) == 0 {
			continue
		}

		suffix := strings.TrimSpace(fileOpts.GetSuffix())
		if len(suffix) == 0 {
			suffix = ".ext.i18n"
		}

		g := plugin.NewGeneratedFile(fmt.Sprintf("%s%s.ts", file.GeneratedFilenamePrefix, suffix),
			file.GoImportPath)
		g.P(pkg.RenderPackageComments(version.Version, "i18n", file.Desc.Path()))
		g.P("namespace ", file.GoPackageName, "I18n {")
		for _, t := range translations {
			g.P()
			g.P("  export const ", t.lang, " = {")
			for _, o := range t.objects {
				g.P("    ", o.name, ": {")
				for _, f := range o.fields {
					g.P(fmt.Sprintf(`      %s: %q,`, f[0], f[1]))
				}
				g.P("    },")
			}
			g.P("  }")
		}
		g.P("}")
	}
	return nil
}

func translationsForEnums(objectsByLang map[string]map[string][][2]string, es []*protogen.Enum) {
	for _, e := range es {
		objectName := e.GoIdent.GoName
		for _, v := range e.Values {
			opts := pkg.ProtoGetExtension[i18n.I18N](v.Desc.Options(), i18n.E_Enum)
			if opts == nil {
				continue
			}
			fieldName := string(v.Desc.Name())
			opts.ProtoReflect().Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				lang := string(d.Name())
				fieldsByObjectName, ok := objectsByLang[lang]
				if !ok {
					fieldsByObjectName = make(map[string][][2]string)
					objectsByLang[lang] = fieldsByObjectName
				}
				fieldsByObjectName[objectName] = append(fieldsByObjectName[objectName], [2]string{fieldName, v.String()})
				return true
			})
		}
	}
}

// map[string]map[string][][2]string
// map[LANG]map[OBJECT][]{FIELD,TEXT}
func translationsForMessages(objectsByLang map[string]map[string][][2]string, ms []*protogen.Message) {
	for _, m := range ms {
		translationsForEnums(objectsByLang, m.Enums)

		objectName := m.GoIdent.GoName
		for _, f := range m.Fields {
			opts := pkg.ProtoGetExtension[i18n.I18N](f.Desc.Options(), i18n.E_Field)
			if opts == nil {
				continue
			}
			fieldName := string(f.Desc.Name())
			opts.ProtoReflect().Range(func(d protoreflect.FieldDescriptor, v protoreflect.Value) bool {
				lang := string(d.Name())
				fieldsByObjectName, ok := objectsByLang[lang]
				if !ok {
					fieldsByObjectName = make(map[string][][2]string)
					objectsByLang[lang] = fieldsByObjectName
				}
				fieldsByObjectName[objectName] = append(fieldsByObjectName[objectName], [2]string{fieldName, v.String()})
				return true
			})
		}

		translationsForMessages(objectsByLang, m.Messages)
	}
}

type object struct {
	name   string
	fields [][2]string
}

type translation struct {
	lang    string
	objects []object
}
