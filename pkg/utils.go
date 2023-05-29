package pkg

import (
	"fmt"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ProtoGetExtension[T any](m proto.Message, t protoreflect.ExtensionType) *T {
	if !proto.HasExtension(m, t) {
		return nil
	}
	v, ok := proto.GetExtension(m, t).(*T)
	if ok {
		return v
	}
	return nil
}

func RenderPackageComments(version, plugin, source, name string) string {
	return fmt.Sprintf(`// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// version: %s
// plugin: %s
// source: %s

package %s
`, version, plugin, source, name)
}
