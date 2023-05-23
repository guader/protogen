package pkg

import (
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
