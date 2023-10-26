// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: conf/v1/kratos_conf_client.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 客户端
type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rest *Client_REST `protobuf:"bytes,1,opt,name=rest,proto3" json:"rest,omitempty"` // REST服务
	Grpc *Client_GRPC `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"` // gRPC服务
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_kratos_conf_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_client_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetRest() *Client_REST {
	if x != nil {
		return x.Rest
	}
	return nil
}

func (x *Client) GetGrpc() *Client_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

// REST
type Client_REST struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout    *durationpb.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"` // 超时时间
	Middleware *Middleware          `protobuf:"bytes,2,opt,name=middleware,proto3" json:"middleware,omitempty"`
}

func (x *Client_REST) Reset() {
	*x = Client_REST{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_kratos_conf_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client_REST) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_REST) ProtoMessage() {}

func (x *Client_REST) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_REST.ProtoReflect.Descriptor instead.
func (*Client_REST) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_client_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Client_REST) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Client_REST) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

// gPRC
type Client_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout    *durationpb.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"` // 超时时间
	Middleware *Middleware          `protobuf:"bytes,2,opt,name=middleware,proto3" json:"middleware,omitempty"`
}

func (x *Client_GRPC) Reset() {
	*x = Client_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_kratos_conf_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client_GRPC) ProtoMessage() {}

func (x *Client_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_kratos_conf_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client_GRPC.ProtoReflect.Descriptor instead.
func (*Client_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_v1_kratos_conf_client_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Client_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Client_GRPC) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

var File_conf_v1_kratos_conf_client_proto protoreflect.FileDescriptor

var file_conf_v1_kratos_conf_client_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6d, 0x69,
	0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb4,
	0x02, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x72, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x45, 0x53, 0x54, 0x52, 0x04, 0x72, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x52, 0x50,
	0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x6d, 0x0a, 0x04, 0x52, 0x45, 0x53, 0x54, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x1a, 0x6d, 0x0a, 0x04, 0x47, 0x52, 0x50, 0x43, 0x12, 0x33,
	0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x30, 0x0a, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x4d,
	0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64, 0x6c,
	0x65, 0x77, 0x61, 0x72, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x78, 0x37, 0x64, 0x6f, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2d, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_kratos_conf_client_proto_rawDescOnce sync.Once
	file_conf_v1_kratos_conf_client_proto_rawDescData = file_conf_v1_kratos_conf_client_proto_rawDesc
)

func file_conf_v1_kratos_conf_client_proto_rawDescGZIP() []byte {
	file_conf_v1_kratos_conf_client_proto_rawDescOnce.Do(func() {
		file_conf_v1_kratos_conf_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_kratos_conf_client_proto_rawDescData)
	})
	return file_conf_v1_kratos_conf_client_proto_rawDescData
}

var file_conf_v1_kratos_conf_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conf_v1_kratos_conf_client_proto_goTypes = []interface{}{
	(*Client)(nil),              // 0: conf.Client
	(*Client_REST)(nil),         // 1: conf.Client.REST
	(*Client_GRPC)(nil),         // 2: conf.Client.GRPC
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*Middleware)(nil),          // 4: conf.Middleware
}
var file_conf_v1_kratos_conf_client_proto_depIdxs = []int32{
	1, // 0: conf.Client.rest:type_name -> conf.Client.REST
	2, // 1: conf.Client.grpc:type_name -> conf.Client.GRPC
	3, // 2: conf.Client.REST.timeout:type_name -> google.protobuf.Duration
	4, // 3: conf.Client.REST.middleware:type_name -> conf.Middleware
	3, // 4: conf.Client.GRPC.timeout:type_name -> google.protobuf.Duration
	4, // 5: conf.Client.GRPC.middleware:type_name -> conf.Middleware
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_conf_v1_kratos_conf_client_proto_init() }
func file_conf_v1_kratos_conf_client_proto_init() {
	if File_conf_v1_kratos_conf_client_proto != nil {
		return
	}
	file_conf_v1_kratos_conf_middleware_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_v1_kratos_conf_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_conf_v1_kratos_conf_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client_REST); i {
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
		file_conf_v1_kratos_conf_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client_GRPC); i {
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
			RawDescriptor: file_conf_v1_kratos_conf_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_kratos_conf_client_proto_goTypes,
		DependencyIndexes: file_conf_v1_kratos_conf_client_proto_depIdxs,
		MessageInfos:      file_conf_v1_kratos_conf_client_proto_msgTypes,
	}.Build()
	File_conf_v1_kratos_conf_client_proto = out.File
	file_conf_v1_kratos_conf_client_proto_rawDesc = nil
	file_conf_v1_kratos_conf_client_proto_goTypes = nil
	file_conf_v1_kratos_conf_client_proto_depIdxs = nil
}