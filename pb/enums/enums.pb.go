// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.3
// source: enums/enums.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable *bool   `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
	Suffix *string `protobuf:"bytes,2,opt,name=suffix" json:"suffix,omitempty"`
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_enums_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_enums_enums_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileOptions.ProtoReflect.Descriptor instead.
func (*FileOptions) Descriptor() ([]byte, []int) {
	return file_enums_enums_proto_rawDescGZIP(), []int{0}
}

func (x *FileOptions) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return false
}

func (x *FileOptions) GetSuffix() string {
	if x != nil && x.Suffix != nil {
		return *x.Suffix
	}
	return ""
}

type EnumOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrMethod *EnumOptions_ErrMethod `protobuf:"bytes,1,opt,name=errMethod" json:"errMethod,omitempty"`
}

func (x *EnumOptions) Reset() {
	*x = EnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_enums_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumOptions) ProtoMessage() {}

func (x *EnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_enums_enums_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumOptions.ProtoReflect.Descriptor instead.
func (*EnumOptions) Descriptor() ([]byte, []int) {
	return file_enums_enums_proto_rawDescGZIP(), []int{1}
}

func (x *EnumOptions) GetErrMethod() *EnumOptions_ErrMethod {
	if x != nil {
		return x.ErrMethod
	}
	return nil
}

type EnumValueOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg *string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (x *EnumValueOptions) Reset() {
	*x = EnumValueOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_enums_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumValueOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumValueOptions) ProtoMessage() {}

func (x *EnumValueOptions) ProtoReflect() protoreflect.Message {
	mi := &file_enums_enums_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumValueOptions.ProtoReflect.Descriptor instead.
func (*EnumValueOptions) Descriptor() ([]byte, []int) {
	return file_enums_enums_proto_rawDescGZIP(), []int{2}
}

func (x *EnumValueOptions) GetMsg() string {
	if x != nil && x.Msg != nil {
		return *x.Msg
	}
	return ""
}

type EnumOptions_ErrMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable   *bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
	WithGrpc *bool `protobuf:"varint,2,opt,name=withGrpc" json:"withGrpc,omitempty"`
}

func (x *EnumOptions_ErrMethod) Reset() {
	*x = EnumOptions_ErrMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_enums_enums_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumOptions_ErrMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumOptions_ErrMethod) ProtoMessage() {}

func (x *EnumOptions_ErrMethod) ProtoReflect() protoreflect.Message {
	mi := &file_enums_enums_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumOptions_ErrMethod.ProtoReflect.Descriptor instead.
func (*EnumOptions_ErrMethod) Descriptor() ([]byte, []int) {
	return file_enums_enums_proto_rawDescGZIP(), []int{1, 0}
}

func (x *EnumOptions_ErrMethod) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return false
}

func (x *EnumOptions_ErrMethod) GetWithGrpc() bool {
	if x != nil && x.WithGrpc != nil {
		return *x.WithGrpc
	}
	return false
}

var file_enums_enums_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         1335,
		Name:          "enums.fileOptions",
		Tag:           "bytes,1335,opt,name=fileOptions",
		Filename:      "enums/enums.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*EnumOptions)(nil),
		Field:         1335,
		Name:          "enums.enumOptions",
		Tag:           "bytes,1335,opt,name=enumOptions",
		Filename:      "enums/enums.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*EnumValueOptions)(nil),
		Field:         1335,
		Name:          "enums.enum",
		Tag:           "bytes,1335,opt,name=enum",
		Filename:      "enums/enums.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional enums.FileOptions fileOptions = 1335;
	E_FileOptions = &file_enums_enums_proto_extTypes[0] // Extension number not registered yet.
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional enums.EnumOptions enumOptions = 1335;
	E_EnumOptions = &file_enums_enums_proto_extTypes[1] // Extension number not registered yet.
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional enums.EnumValueOptions enum = 1335;
	E_Enum = &file_enums_enums_proto_extTypes[2] // Extension number not registered yet.
)

var File_enums_enums_proto protoreflect.FileDescriptor

var file_enums_enums_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x22, 0x8a, 0x01, 0x0a, 0x0b,
	0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x65,
	0x72, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x45, 0x72, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x09, 0x65, 0x72,
	0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x1a, 0x3f, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x69, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x77, 0x69, 0x74, 0x68, 0x47, 0x72, 0x70, 0x63, 0x22, 0x24, 0x0a, 0x10, 0x45, 0x6e, 0x75, 0x6d,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x3a, 0x53,
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x3a, 0x53, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xb7, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x65, 0x6e, 0x75,
	0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x4f, 0x0a, 0x04, 0x65, 0x6e, 0x75, 0x6d,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb7, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75, 0x61, 0x64, 0x65, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
}

var (
	file_enums_enums_proto_rawDescOnce sync.Once
	file_enums_enums_proto_rawDescData = file_enums_enums_proto_rawDesc
)

func file_enums_enums_proto_rawDescGZIP() []byte {
	file_enums_enums_proto_rawDescOnce.Do(func() {
		file_enums_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_enums_enums_proto_rawDescData)
	})
	return file_enums_enums_proto_rawDescData
}

var file_enums_enums_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_enums_enums_proto_goTypes = []interface{}{
	(*FileOptions)(nil),                   // 0: enums.FileOptions
	(*EnumOptions)(nil),                   // 1: enums.EnumOptions
	(*EnumValueOptions)(nil),              // 2: enums.EnumValueOptions
	(*EnumOptions_ErrMethod)(nil),         // 3: enums.EnumOptions.ErrMethod
	(*descriptorpb.FileOptions)(nil),      // 4: google.protobuf.FileOptions
	(*descriptorpb.EnumOptions)(nil),      // 5: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 6: google.protobuf.EnumValueOptions
}
var file_enums_enums_proto_depIdxs = []int32{
	3, // 0: enums.EnumOptions.errMethod:type_name -> enums.EnumOptions.ErrMethod
	4, // 1: enums.fileOptions:extendee -> google.protobuf.FileOptions
	5, // 2: enums.enumOptions:extendee -> google.protobuf.EnumOptions
	6, // 3: enums.enum:extendee -> google.protobuf.EnumValueOptions
	0, // 4: enums.fileOptions:type_name -> enums.FileOptions
	1, // 5: enums.enumOptions:type_name -> enums.EnumOptions
	2, // 6: enums.enum:type_name -> enums.EnumValueOptions
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	4, // [4:7] is the sub-list for extension type_name
	1, // [1:4] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_enums_enums_proto_init() }
func file_enums_enums_proto_init() {
	if File_enums_enums_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_enums_enums_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enums_enums_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enums_enums_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumValueOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_enums_enums_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumOptions_ErrMethod); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_enums_enums_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_enums_enums_proto_goTypes,
		DependencyIndexes: file_enums_enums_proto_depIdxs,
		MessageInfos:      file_enums_enums_proto_msgTypes,
		ExtensionInfos:    file_enums_enums_proto_extTypes,
	}.Build()
	File_enums_enums_proto = out.File
	file_enums_enums_proto_rawDesc = nil
	file_enums_enums_proto_goTypes = nil
	file_enums_enums_proto_depIdxs = nil
}
