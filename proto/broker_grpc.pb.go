// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: broker.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BrokerApiClient is the client API for BrokerApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerApiClient interface {
	// returns a stream of transactions
	StreamReceivedTransactions(ctx context.Context, in *TxStreamRequest, opts ...grpc.CallOption) (BrokerApi_StreamReceivedTransactionsClient, error)
}

type brokerApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerApiClient(cc grpc.ClientConnInterface) BrokerApiClient {
	return &brokerApiClient{cc}
}

func (c *brokerApiClient) StreamReceivedTransactions(ctx context.Context, in *TxStreamRequest, opts ...grpc.CallOption) (BrokerApi_StreamReceivedTransactionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &BrokerApi_ServiceDesc.Streams[0], "/usemerkle.com.broker.proto.BrokerApi/StreamReceivedTransactions", opts...)
	if err != nil {
		return nil, err
	}
	x := &brokerApiStreamReceivedTransactionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BrokerApi_StreamReceivedTransactionsClient interface {
	Recv() (*Transaction, error)
	grpc.ClientStream
}

type brokerApiStreamReceivedTransactionsClient struct {
	grpc.ClientStream
}

func (x *brokerApiStreamReceivedTransactionsClient) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BrokerApiServer is the server API for BrokerApi service.
// All implementations must embed UnimplementedBrokerApiServer
// for forward compatibility
type BrokerApiServer interface {
	// returns a stream of transactions
	StreamReceivedTransactions(*TxStreamRequest, BrokerApi_StreamReceivedTransactionsServer) error
	mustEmbedUnimplementedBrokerApiServer()
}

// UnimplementedBrokerApiServer must be embedded to have forward compatible implementations.
type UnimplementedBrokerApiServer struct {
}

func (UnimplementedBrokerApiServer) StreamReceivedTransactions(*TxStreamRequest, BrokerApi_StreamReceivedTransactionsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamReceivedTransactions not implemented")
}
func (UnimplementedBrokerApiServer) mustEmbedUnimplementedBrokerApiServer() {}

// UnsafeBrokerApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerApiServer will
// result in compilation errors.
type UnsafeBrokerApiServer interface {
	mustEmbedUnimplementedBrokerApiServer()
}

func RegisterBrokerApiServer(s grpc.ServiceRegistrar, srv BrokerApiServer) {
	s.RegisterService(&BrokerApi_ServiceDesc, srv)
}

func _BrokerApi_StreamReceivedTransactions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TxStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BrokerApiServer).StreamReceivedTransactions(m, &brokerApiStreamReceivedTransactionsServer{stream})
}

type BrokerApi_StreamReceivedTransactionsServer interface {
	Send(*Transaction) error
	grpc.ServerStream
}

type brokerApiStreamReceivedTransactionsServer struct {
	grpc.ServerStream
}

func (x *brokerApiStreamReceivedTransactionsServer) Send(m *Transaction) error {
	return x.ServerStream.SendMsg(m)
}

// BrokerApi_ServiceDesc is the grpc.ServiceDesc for BrokerApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usemerkle.com.broker.proto.BrokerApi",
	HandlerType: (*BrokerApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamReceivedTransactions",
			Handler:       _BrokerApi_StreamReceivedTransactions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "broker.proto",
}
