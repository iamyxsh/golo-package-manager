// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: pb/file.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LinkMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Package string `protobuf:"bytes,1,opt,name=package,proto3" json:"package,omitempty"`
}

func (x *LinkMessageRequest) Reset() {
	*x = LinkMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkMessageRequest) ProtoMessage() {}

func (x *LinkMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkMessageRequest.ProtoReflect.Descriptor instead.
func (*LinkMessageRequest) Descriptor() ([]byte, []int) {
	return file_pb_file_proto_rawDescGZIP(), []int{0}
}

func (x *LinkMessageRequest) GetPackage() string {
	if x != nil {
		return x.Package
	}
	return ""
}

type LinkMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
}

func (x *LinkMessageResponse) Reset() {
	*x = LinkMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LinkMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkMessageResponse) ProtoMessage() {}

func (x *LinkMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkMessageResponse.ProtoReflect.Descriptor instead.
func (*LinkMessageResponse) Descriptor() ([]byte, []int) {
	return file_pb_file_proto_rawDescGZIP(), []int{1}
}

func (x *LinkMessageResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_pb_file_proto protoreflect.FileDescriptor

var file_pb_file_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x2e, 0x0a, 0x12, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x32, 0x49,
	0x0a, 0x0b, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_file_proto_rawDescOnce sync.Once
	file_pb_file_proto_rawDescData = file_pb_file_proto_rawDesc
)

func file_pb_file_proto_rawDescGZIP() []byte {
	file_pb_file_proto_rawDescOnce.Do(func() {
		file_pb_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_file_proto_rawDescData)
	})
	return file_pb_file_proto_rawDescData
}

var file_pb_file_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_file_proto_goTypes = []interface{}{
	(*LinkMessageRequest)(nil),  // 0: pb.LinkMessageRequest
	(*LinkMessageResponse)(nil), // 1: pb.LinkMessageResponse
}
var file_pb_file_proto_depIdxs = []int32{
	0, // 0: pb.LinkService.GetLink:input_type -> pb.LinkMessageRequest
	1, // 1: pb.LinkService.GetLink:output_type -> pb.LinkMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_file_proto_init() }
func file_pb_file_proto_init() {
	if File_pb_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_file_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkMessageRequest); i {
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
		file_pb_file_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LinkMessageResponse); i {
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
			RawDescriptor: file_pb_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_file_proto_goTypes,
		DependencyIndexes: file_pb_file_proto_depIdxs,
		MessageInfos:      file_pb_file_proto_msgTypes,
	}.Build()
	File_pb_file_proto = out.File
	file_pb_file_proto_rawDesc = nil
	file_pb_file_proto_goTypes = nil
	file_pb_file_proto_depIdxs = nil
}
