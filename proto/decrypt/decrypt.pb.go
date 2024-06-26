// protoc --go_out=. --go-grpc_out=. decrypt.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: decrypt.proto

package decrypt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DecryptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *structpb.Struct `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *DecryptRequest) Reset() {
	*x = DecryptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decrypt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptRequest) ProtoMessage() {}

func (x *DecryptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_decrypt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptRequest.ProtoReflect.Descriptor instead.
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return file_decrypt_proto_rawDescGZIP(), []int{0}
}

func (x *DecryptRequest) GetRequest() *structpb.Struct {
	if x != nil {
		return x.Request
	}
	return nil
}

type DecryptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *structpb.Struct `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DecryptResponse) Reset() {
	*x = DecryptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_decrypt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecryptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecryptResponse) ProtoMessage() {}

func (x *DecryptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_decrypt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecryptResponse.ProtoReflect.Descriptor instead.
func (*DecryptResponse) Descriptor() ([]byte, []int) {
	return file_decrypt_proto_rawDescGZIP(), []int{1}
}

func (x *DecryptResponse) GetResponse() *structpb.Struct {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_decrypt_proto protoreflect.FileDescriptor

var file_decrypt_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a, 0x0f, 0x44,
	0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0x49, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x3e,
	0x0a, 0x07, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x2e, 0x44, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b,
	0x5a, 0x09, 0x2e, 0x2f, 0x64, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_decrypt_proto_rawDescOnce sync.Once
	file_decrypt_proto_rawDescData = file_decrypt_proto_rawDesc
)

func file_decrypt_proto_rawDescGZIP() []byte {
	file_decrypt_proto_rawDescOnce.Do(func() {
		file_decrypt_proto_rawDescData = protoimpl.X.CompressGZIP(file_decrypt_proto_rawDescData)
	})
	return file_decrypt_proto_rawDescData
}

var file_decrypt_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_decrypt_proto_goTypes = []any{
	(*DecryptRequest)(nil),  // 0: decrypt.DecryptRequest
	(*DecryptResponse)(nil), // 1: decrypt.DecryptResponse
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
}
var file_decrypt_proto_depIdxs = []int32{
	2, // 0: decrypt.DecryptRequest.request:type_name -> google.protobuf.Struct
	2, // 1: decrypt.DecryptResponse.response:type_name -> google.protobuf.Struct
	0, // 2: decrypt.Decrypt.Decrypt:input_type -> decrypt.DecryptRequest
	1, // 3: decrypt.Decrypt.Decrypt:output_type -> decrypt.DecryptResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_decrypt_proto_init() }
func file_decrypt_proto_init() {
	if File_decrypt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_decrypt_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DecryptRequest); i {
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
		file_decrypt_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DecryptResponse); i {
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
			RawDescriptor: file_decrypt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_decrypt_proto_goTypes,
		DependencyIndexes: file_decrypt_proto_depIdxs,
		MessageInfos:      file_decrypt_proto_msgTypes,
	}.Build()
	File_decrypt_proto = out.File
	file_decrypt_proto_rawDesc = nil
	file_decrypt_proto_goTypes = nil
	file_decrypt_proto_depIdxs = nil
}
