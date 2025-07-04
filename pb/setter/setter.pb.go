// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: setter/setter.proto

package setter

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileOptions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enable        *bool                  `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
	Suffix        *string                `protobuf:"bytes,2,opt,name=suffix" json:"suffix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FileOptions) Reset() {
	*x = FileOptions{}
	mi := &file_setter_setter_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileOptions) ProtoMessage() {}

func (x *FileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_setter_setter_proto_msgTypes[0]
	if x != nil {
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
	return file_setter_setter_proto_rawDescGZIP(), []int{0}
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

type MessageOptions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enable        *bool                  `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	mi := &file_setter_setter_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_setter_setter_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_setter_setter_proto_rawDescGZIP(), []int{1}
}

func (x *MessageOptions) GetEnable() bool {
	if x != nil && x.Enable != nil {
		return *x.Enable
	}
	return false
}

type FieldOptions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          *string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldOptions) Reset() {
	*x = FieldOptions{}
	mi := &file_setter_setter_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldOptions) ProtoMessage() {}

func (x *FieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_setter_setter_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldOptions.ProtoReflect.Descriptor instead.
func (*FieldOptions) Descriptor() ([]byte, []int) {
	return file_setter_setter_proto_rawDescGZIP(), []int{2}
}

func (x *FieldOptions) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

var file_setter_setter_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileOptions)(nil),
		Field:         1330,
		Name:          "setter.fileOptions",
		Tag:           "bytes,1330,opt,name=fileOptions",
		Filename:      "setter/setter.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         1330,
		Name:          "setter.messageOptions",
		Tag:           "bytes,1330,opt,name=messageOptions",
		Filename:      "setter/setter.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*FieldOptions)(nil),
		Field:         1330,
		Name:          "setter.fieldOptions",
		Tag:           "bytes,1330,opt,name=fieldOptions",
		Filename:      "setter/setter.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional setter.FileOptions fileOptions = 1330;
	E_FileOptions = &file_setter_setter_proto_extTypes[0] // Extension number not registered yet.
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional setter.MessageOptions messageOptions = 1330;
	E_MessageOptions = &file_setter_setter_proto_extTypes[1] // Extension number not registered yet.
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional setter.FieldOptions fieldOptions = 1330;
	E_FieldOptions = &file_setter_setter_proto_extTypes[2] // Extension number not registered yet.
)

var File_setter_setter_proto protoreflect.FileDescriptor

const file_setter_setter_proto_rawDesc = "" +
	"\n" +
	"\x13setter/setter.proto\x12\x06setter\x1a google/protobuf/descriptor.proto\"=\n" +
	"\vFileOptions\x12\x16\n" +
	"\x06enable\x18\x01 \x01(\bR\x06enable\x12\x16\n" +
	"\x06suffix\x18\x02 \x01(\tR\x06suffix\"(\n" +
	"\x0eMessageOptions\x12\x16\n" +
	"\x06enable\x18\x01 \x01(\bR\x06enable\"\"\n" +
	"\fFieldOptions\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name:T\n" +
	"\vfileOptions\x12\x1c.google.protobuf.FileOptions\x18\xb2\n" +
	" \x01(\v2\x13.setter.FileOptionsR\vfileOptions:`\n" +
	"\x0emessageOptions\x12\x1f.google.protobuf.MessageOptions\x18\xb2\n" +
	" \x01(\v2\x16.setter.MessageOptionsR\x0emessageOptions:X\n" +
	"\ffieldOptions\x12\x1d.google.protobuf.FieldOptions\x18\xb2\n" +
	" \x01(\v2\x14.setter.FieldOptionsR\ffieldOptionsB&Z$github.com/guader/protogen/pb/setter"

var (
	file_setter_setter_proto_rawDescOnce sync.Once
	file_setter_setter_proto_rawDescData []byte
)

func file_setter_setter_proto_rawDescGZIP() []byte {
	file_setter_setter_proto_rawDescOnce.Do(func() {
		file_setter_setter_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_setter_setter_proto_rawDesc), len(file_setter_setter_proto_rawDesc)))
	})
	return file_setter_setter_proto_rawDescData
}

var file_setter_setter_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_setter_setter_proto_goTypes = []any{
	(*FileOptions)(nil),                 // 0: setter.FileOptions
	(*MessageOptions)(nil),              // 1: setter.MessageOptions
	(*FieldOptions)(nil),                // 2: setter.FieldOptions
	(*descriptorpb.FileOptions)(nil),    // 3: google.protobuf.FileOptions
	(*descriptorpb.MessageOptions)(nil), // 4: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 5: google.protobuf.FieldOptions
}
var file_setter_setter_proto_depIdxs = []int32{
	3, // 0: setter.fileOptions:extendee -> google.protobuf.FileOptions
	4, // 1: setter.messageOptions:extendee -> google.protobuf.MessageOptions
	5, // 2: setter.fieldOptions:extendee -> google.protobuf.FieldOptions
	0, // 3: setter.fileOptions:type_name -> setter.FileOptions
	1, // 4: setter.messageOptions:type_name -> setter.MessageOptions
	2, // 5: setter.fieldOptions:type_name -> setter.FieldOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	3, // [3:6] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_setter_setter_proto_init() }
func file_setter_setter_proto_init() {
	if File_setter_setter_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_setter_setter_proto_rawDesc), len(file_setter_setter_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_setter_setter_proto_goTypes,
		DependencyIndexes: file_setter_setter_proto_depIdxs,
		MessageInfos:      file_setter_setter_proto_msgTypes,
		ExtensionInfos:    file_setter_setter_proto_extTypes,
	}.Build()
	File_setter_setter_proto = out.File
	file_setter_setter_proto_goTypes = nil
	file_setter_setter_proto_depIdxs = nil
}
