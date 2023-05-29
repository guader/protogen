// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: errorer/errorer.proto

package errorer

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
		mi := &file_errorer_errorer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_errorer_errorer_proto_msgTypes[0]
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
	return file_errorer_errorer_proto_rawDescGZIP(), []int{0}
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

	Enable *bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (x *EnumOptions) Reset() {
	*x = EnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorer_errorer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumOptions) ProtoMessage() {}

func (x *EnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_errorer_errorer_proto_msgTypes[1]
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
	return file_errorer_errorer_proto_rawDescGZIP(), []int{1}
}

func (x *EnumOptions) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return false
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_errorer_errorer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_errorer_errorer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_errorer_errorer_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var file_errorer_errorer_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         1332,
		Name:          "errorer.fileOptions",
		Tag:           "bytes,1332,opt,name=fileOptions",
		Filename:      "errorer/errorer.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*EnumOptions)(nil),
		Field:         1332,
		Name:          "errorer.enumOptions",
		Tag:           "bytes,1332,opt,name=enumOptions",
		Filename:      "errorer/errorer.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*Error)(nil),
		Field:         1332,
		Name:          "errorer.err",
		Tag:           "bytes,1332,opt,name=err",
		Filename:      "errorer/errorer.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional errorer.FileOptions fileOptions = 1332;
	E_FileOptions = &file_errorer_errorer_proto_extTypes[0] // Extension number not registered yet.
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional errorer.EnumOptions enumOptions = 1332;
	E_EnumOptions = &file_errorer_errorer_proto_extTypes[1] // Extension number not registered yet.
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional errorer.Error err = 1332;
	E_Err = &file_errorer_errorer_proto_extTypes[2] // Extension number not registered yet.
)

var File_errorer_errorer_proto protoreflect.FileDescriptor

var file_errorer_errorer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x72,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66,
	0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x22, 0x25, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x55, 0x0a, 0x0b, 0x66,
	0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb4, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x55, 0x0a, 0x0b, 0x65, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xb4, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x72,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x65, 0x6e,
	0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x44, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xb4, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x65, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x75,
	0x61, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x65, 0x72,
}

var (
	file_errorer_errorer_proto_rawDescOnce sync.Once
	file_errorer_errorer_proto_rawDescData = file_errorer_errorer_proto_rawDesc
)

func file_errorer_errorer_proto_rawDescGZIP() []byte {
	file_errorer_errorer_proto_rawDescOnce.Do(func() {
		file_errorer_errorer_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorer_errorer_proto_rawDescData)
	})
	return file_errorer_errorer_proto_rawDescData
}

var file_errorer_errorer_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_errorer_errorer_proto_goTypes = []interface{}{
	(*FileOptions)(nil),                   // 0: errorer.FileOptions
	(*EnumOptions)(nil),                   // 1: errorer.EnumOptions
	(*Error)(nil),                         // 2: errorer.Error
	(*descriptorpb.FileOptions)(nil),      // 3: google.protobuf.FileOptions
	(*descriptorpb.EnumOptions)(nil),      // 4: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 5: google.protobuf.EnumValueOptions
}
var file_errorer_errorer_proto_depIdxs = []int32{
	3, // 0: errorer.fileOptions:extendee -> google.protobuf.FileOptions
	4, // 1: errorer.enumOptions:extendee -> google.protobuf.EnumOptions
	5, // 2: errorer.err:extendee -> google.protobuf.EnumValueOptions
	0, // 3: errorer.fileOptions:type_name -> errorer.FileOptions
	1, // 4: errorer.enumOptions:type_name -> errorer.EnumOptions
	2, // 5: errorer.err:type_name -> errorer.Error
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errorer_errorer_proto_init() }
func file_errorer_errorer_proto_init() {
	if File_errorer_errorer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_errorer_errorer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_errorer_errorer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_errorer_errorer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_errorer_errorer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_errorer_errorer_proto_goTypes,
		DependencyIndexes: file_errorer_errorer_proto_depIdxs,
		MessageInfos:      file_errorer_errorer_proto_msgTypes,
		ExtensionInfos:    file_errorer_errorer_proto_extTypes,
	}.Build()
	File_errorer_errorer_proto = out.File
	file_errorer_errorer_proto_rawDesc = nil
	file_errorer_errorer_proto_goTypes = nil
	file_errorer_errorer_proto_depIdxs = nil
}
