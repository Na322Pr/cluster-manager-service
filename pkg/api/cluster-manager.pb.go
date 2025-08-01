// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/cluster-manager.proto

package cluster_manager_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SetNodeCountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	NodeCount     int64                  `protobuf:"varint,1,opt,name=node_count,json=nodeCount,proto3" json:"node_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNodeCountRequest) Reset() {
	*x = SetNodeCountRequest{}
	mi := &file_api_cluster_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNodeCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNodeCountRequest) ProtoMessage() {}

func (x *SetNodeCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNodeCountRequest.ProtoReflect.Descriptor instead.
func (*SetNodeCountRequest) Descriptor() ([]byte, []int) {
	return file_api_cluster_manager_proto_rawDescGZIP(), []int{0}
}

func (x *SetNodeCountRequest) GetNodeCount() int64 {
	if x != nil {
		return x.NodeCount
	}
	return 0
}

type SetNodeCountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        bool                   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetNodeCountResponse) Reset() {
	*x = SetNodeCountResponse{}
	mi := &file_api_cluster_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetNodeCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNodeCountResponse) ProtoMessage() {}

func (x *SetNodeCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_cluster_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNodeCountResponse.ProtoReflect.Descriptor instead.
func (*SetNodeCountResponse) Descriptor() ([]byte, []int) {
	return file_api_cluster_manager_proto_rawDescGZIP(), []int{1}
}

func (x *SetNodeCountResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_api_cluster_manager_proto protoreflect.FileDescriptor

const file_api_cluster_manager_proto_rawDesc = "" +
	"\n" +
	"\x19api/cluster-manager.proto\x12\x17cluster_manager_service\"4\n" +
	"\x13SetNodeCountRequest\x12\x1d\n" +
	"\n" +
	"node_count\x18\x01 \x01(\x03R\tnodeCount\".\n" +
	"\x14SetNodeCountResponse\x12\x16\n" +
	"\x06result\x18\x01 \x01(\bR\x06result2}\n" +
	"\x0eClusterManager\x12k\n" +
	"\fSetNodeCount\x12,.cluster_manager_service.SetNodeCountRequest\x1a-.cluster_manager_service.SetNodeCountResponseB`Z^github.com/Na322Pr/cluster-manager-service/pkg/cluster-manager-service;cluster_manager_serviceb\x06proto3"

var (
	file_api_cluster_manager_proto_rawDescOnce sync.Once
	file_api_cluster_manager_proto_rawDescData []byte
)

func file_api_cluster_manager_proto_rawDescGZIP() []byte {
	file_api_cluster_manager_proto_rawDescOnce.Do(func() {
		file_api_cluster_manager_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_cluster_manager_proto_rawDesc), len(file_api_cluster_manager_proto_rawDesc)))
	})
	return file_api_cluster_manager_proto_rawDescData
}

var file_api_cluster_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_cluster_manager_proto_goTypes = []any{
	(*SetNodeCountRequest)(nil),  // 0: cluster_manager_service.SetNodeCountRequest
	(*SetNodeCountResponse)(nil), // 1: cluster_manager_service.SetNodeCountResponse
}
var file_api_cluster_manager_proto_depIdxs = []int32{
	0, // 0: cluster_manager_service.ClusterManager.SetNodeCount:input_type -> cluster_manager_service.SetNodeCountRequest
	1, // 1: cluster_manager_service.ClusterManager.SetNodeCount:output_type -> cluster_manager_service.SetNodeCountResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_cluster_manager_proto_init() }
func file_api_cluster_manager_proto_init() {
	if File_api_cluster_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_cluster_manager_proto_rawDesc), len(file_api_cluster_manager_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_cluster_manager_proto_goTypes,
		DependencyIndexes: file_api_cluster_manager_proto_depIdxs,
		MessageInfos:      file_api_cluster_manager_proto_msgTypes,
	}.Build()
	File_api_cluster_manager_proto = out.File
	file_api_cluster_manager_proto_goTypes = nil
	file_api_cluster_manager_proto_depIdxs = nil
}
