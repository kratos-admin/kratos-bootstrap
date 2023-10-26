// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: conf/v1/config.proto

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

// 配置服务
type RemoteConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string                   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Etcd       *RemoteConfig_Etcd       `protobuf:"bytes,2,opt,name=etcd,proto3" json:"etcd,omitempty"`
	Consul     *RemoteConfig_Consul     `protobuf:"bytes,3,opt,name=consul,proto3" json:"consul,omitempty"`
	Nacos      *RemoteConfig_Nacos      `protobuf:"bytes,4,opt,name=nacos,proto3" json:"nacos,omitempty"`
	Apollo     *RemoteConfig_Apollo     `protobuf:"bytes,6,opt,name=apollo,proto3" json:"apollo,omitempty"`
	Kubernetes *RemoteConfig_Kubernetes `protobuf:"bytes,7,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Polaris    *RemoteConfig_Polaris    `protobuf:"bytes,8,opt,name=polaris,proto3" json:"polaris,omitempty"`
}

func (x *RemoteConfig) Reset() {
	*x = RemoteConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig) ProtoMessage() {}

func (x *RemoteConfig) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig.ProtoReflect.Descriptor instead.
func (*RemoteConfig) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *RemoteConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RemoteConfig) GetEtcd() *RemoteConfig_Etcd {
	if x != nil {
		return x.Etcd
	}
	return nil
}

func (x *RemoteConfig) GetConsul() *RemoteConfig_Consul {
	if x != nil {
		return x.Consul
	}
	return nil
}

func (x *RemoteConfig) GetNacos() *RemoteConfig_Nacos {
	if x != nil {
		return x.Nacos
	}
	return nil
}

func (x *RemoteConfig) GetApollo() *RemoteConfig_Apollo {
	if x != nil {
		return x.Apollo
	}
	return nil
}

func (x *RemoteConfig) GetKubernetes() *RemoteConfig_Kubernetes {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

func (x *RemoteConfig) GetPolaris() *RemoteConfig_Polaris {
	if x != nil {
		return x.Polaris
	}
	return nil
}

type RemoteConfig_Nacos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"` // 服务端地址
	Port    uint64 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`      // 服务端端口
	Key     string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`         //
}

func (x *RemoteConfig_Nacos) Reset() {
	*x = RemoteConfig_Nacos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Nacos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Nacos) ProtoMessage() {}

func (x *RemoteConfig_Nacos) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Nacos.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Nacos) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RemoteConfig_Nacos) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RemoteConfig_Nacos) GetPort() uint64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RemoteConfig_Nacos) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RemoteConfig_Etcd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints []string             `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Timeout   *durationpb.Duration `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Key       string               `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"` //
}

func (x *RemoteConfig_Etcd) Reset() {
	*x = RemoteConfig_Etcd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Etcd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Etcd) ProtoMessage() {}

func (x *RemoteConfig_Etcd) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Etcd.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Etcd) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *RemoteConfig_Etcd) GetEndpoints() []string {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *RemoteConfig_Etcd) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RemoteConfig_Etcd) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RemoteConfig_Consul struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme  string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`   // 网络样式
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"` // 服务端地址
	Key     string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`         //
}

func (x *RemoteConfig_Consul) Reset() {
	*x = RemoteConfig_Consul{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Consul) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Consul) ProtoMessage() {}

func (x *RemoteConfig_Consul) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Consul.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Consul) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *RemoteConfig_Consul) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *RemoteConfig_Consul) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RemoteConfig_Consul) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RemoteConfig_Apollo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint  string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AppId     string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	Cluster   string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Secret    string `protobuf:"bytes,5,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RemoteConfig_Apollo) Reset() {
	*x = RemoteConfig_Apollo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Apollo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Apollo) ProtoMessage() {}

func (x *RemoteConfig_Apollo) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Apollo.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Apollo) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0, 3}
}

func (x *RemoteConfig_Apollo) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *RemoteConfig_Apollo) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *RemoteConfig_Apollo) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *RemoteConfig_Apollo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *RemoteConfig_Apollo) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type RemoteConfig_Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *RemoteConfig_Kubernetes) Reset() {
	*x = RemoteConfig_Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Kubernetes) ProtoMessage() {}

func (x *RemoteConfig_Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Kubernetes.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Kubernetes) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0, 4}
}

func (x *RemoteConfig_Kubernetes) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type RemoteConfig_Polaris struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoteConfig_Polaris) Reset() {
	*x = RemoteConfig_Polaris{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_v1_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteConfig_Polaris) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteConfig_Polaris) ProtoMessage() {}

func (x *RemoteConfig_Polaris) ProtoReflect() protoreflect.Message {
	mi := &file_conf_v1_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteConfig_Polaris.ProtoReflect.Descriptor instead.
func (*RemoteConfig_Polaris) Descriptor() ([]byte, []int) {
	return file_conf_v1_config_proto_rawDescGZIP(), []int{0, 5}
}

var File_conf_v1_config_proto protoreflect.FileDescriptor

var file_conf_v1_config_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x06, 0x0a,
	0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x2b, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x45, 0x74, 0x63, 0x64, 0x52, 0x04, 0x65, 0x74, 0x63, 0x64, 0x12, 0x31,
	0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x12, 0x2e, 0x0a, 0x05, 0x6e, 0x61, 0x63, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4e, 0x61, 0x63, 0x6f, 0x73, 0x52, 0x05, 0x6e, 0x61, 0x63, 0x6f,
	0x73, 0x12, 0x31, 0x0a, 0x06, 0x61, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x52, 0x06, 0x61, 0x70,
	0x6f, 0x6c, 0x6c, 0x6f, 0x12, 0x3d, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e,
	0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73,
	0x52, 0x07, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x1a, 0x47, 0x0a, 0x05, 0x4e, 0x61, 0x63,
	0x6f, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x1a, 0x6b, 0x0a, 0x04, 0x45, 0x74, 0x63, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a,
	0x4c, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x1a, 0x8b, 0x01,
	0x0a, 0x06, 0x41, 0x70, 0x6f, 0x6c, 0x6c, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x1a, 0x2a, 0x0a, 0x0a, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x1a, 0x09, 0x0a, 0x07, 0x50, 0x6f, 0x6c, 0x61, 0x72,
	0x69, 0x73, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x78, 0x37, 0x64, 0x6f, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_v1_config_proto_rawDescOnce sync.Once
	file_conf_v1_config_proto_rawDescData = file_conf_v1_config_proto_rawDesc
)

func file_conf_v1_config_proto_rawDescGZIP() []byte {
	file_conf_v1_config_proto_rawDescOnce.Do(func() {
		file_conf_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_v1_config_proto_rawDescData)
	})
	return file_conf_v1_config_proto_rawDescData
}

var file_conf_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_conf_v1_config_proto_goTypes = []interface{}{
	(*RemoteConfig)(nil),            // 0: conf.RemoteConfig
	(*RemoteConfig_Nacos)(nil),      // 1: conf.RemoteConfig.Nacos
	(*RemoteConfig_Etcd)(nil),       // 2: conf.RemoteConfig.Etcd
	(*RemoteConfig_Consul)(nil),     // 3: conf.RemoteConfig.Consul
	(*RemoteConfig_Apollo)(nil),     // 4: conf.RemoteConfig.Apollo
	(*RemoteConfig_Kubernetes)(nil), // 5: conf.RemoteConfig.Kubernetes
	(*RemoteConfig_Polaris)(nil),    // 6: conf.RemoteConfig.Polaris
	(*durationpb.Duration)(nil),     // 7: google.protobuf.Duration
}
var file_conf_v1_config_proto_depIdxs = []int32{
	2, // 0: conf.RemoteConfig.etcd:type_name -> conf.RemoteConfig.Etcd
	3, // 1: conf.RemoteConfig.consul:type_name -> conf.RemoteConfig.Consul
	1, // 2: conf.RemoteConfig.nacos:type_name -> conf.RemoteConfig.Nacos
	4, // 3: conf.RemoteConfig.apollo:type_name -> conf.RemoteConfig.Apollo
	5, // 4: conf.RemoteConfig.kubernetes:type_name -> conf.RemoteConfig.Kubernetes
	6, // 5: conf.RemoteConfig.polaris:type_name -> conf.RemoteConfig.Polaris
	7, // 6: conf.RemoteConfig.Etcd.timeout:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_conf_v1_config_proto_init() }
func file_conf_v1_config_proto_init() {
	if File_conf_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig); i {
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
		file_conf_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Nacos); i {
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
		file_conf_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Etcd); i {
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
		file_conf_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Consul); i {
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
		file_conf_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Apollo); i {
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
		file_conf_v1_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Kubernetes); i {
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
		file_conf_v1_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteConfig_Polaris); i {
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
			RawDescriptor: file_conf_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_v1_config_proto_goTypes,
		DependencyIndexes: file_conf_v1_config_proto_depIdxs,
		MessageInfos:      file_conf_v1_config_proto_msgTypes,
	}.Build()
	File_conf_v1_config_proto = out.File
	file_conf_v1_config_proto_rawDesc = nil
	file_conf_v1_config_proto_goTypes = nil
	file_conf_v1_config_proto_depIdxs = nil
}
