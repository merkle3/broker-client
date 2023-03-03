// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: broker.proto

package proto

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash  string `protobuf:"bytes,1,opt,name=txHash,proto3" json:"txHash,omitempty"`
	TxBytes string `protobuf:"bytes,2,opt,name=txBytes,proto3" json:"txBytes,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *Transaction) GetTxBytes() string {
	if x != nil {
		return x.TxBytes
	}
	return ""
}

type TxStreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TxStreamRequest) Reset() {
	*x = TxStreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_broker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxStreamRequest) ProtoMessage() {}

func (x *TxStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_broker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxStreamRequest.ProtoReflect.Descriptor instead.
func (*TxStreamRequest) Descriptor() ([]byte, []int) {
	return file_broker_proto_rawDescGZIP(), []int{1}
}

func (x *TxStreamRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_broker_proto protoreflect.FileDescriptor

var file_broker_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a,
	0x75, 0x73, 0x65, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0b, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x25, 0x0a, 0x0f, 0x54,
	0x78, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x32, 0x83, 0x01, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x41, 0x70, 0x69,
	0x12, 0x76, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b,
	0x2e, 0x75, 0x73, 0x65, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x78, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x75, 0x73,
	0x65, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1c, 0x5a, 0x1a, 0x75, 0x73, 0x65, 0x6d,
	0x65, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_broker_proto_rawDescOnce sync.Once
	file_broker_proto_rawDescData = file_broker_proto_rawDesc
)

func file_broker_proto_rawDescGZIP() []byte {
	file_broker_proto_rawDescOnce.Do(func() {
		file_broker_proto_rawDescData = protoimpl.X.CompressGZIP(file_broker_proto_rawDescData)
	})
	return file_broker_proto_rawDescData
}

var file_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_broker_proto_goTypes = []interface{}{
	(*Transaction)(nil),     // 0: usemerkle.com.broker.proto.Transaction
	(*TxStreamRequest)(nil), // 1: usemerkle.com.broker.proto.TxStreamRequest
}
var file_broker_proto_depIdxs = []int32{
	1, // 0: usemerkle.com.broker.proto.BrokerApi.StreamReceivedTransactions:input_type -> usemerkle.com.broker.proto.TxStreamRequest
	0, // 1: usemerkle.com.broker.proto.BrokerApi.StreamReceivedTransactions:output_type -> usemerkle.com.broker.proto.Transaction
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_broker_proto_init() }
func file_broker_proto_init() {
	if File_broker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_broker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_broker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxStreamRequest); i {
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
			RawDescriptor: file_broker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_broker_proto_goTypes,
		DependencyIndexes: file_broker_proto_depIdxs,
		MessageInfos:      file_broker_proto_msgTypes,
	}.Build()
	File_broker_proto = out.File
	file_broker_proto_rawDesc = nil
	file_broker_proto_goTypes = nil
	file_broker_proto_depIdxs = nil
}
