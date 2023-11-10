// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/intent.proto

package api

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

// IntentControllerClient is the client API for IntentController service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IntentControllerClient interface {
	GetIntentPath(ctx context.Context, opts ...grpc.CallOption) (IntentController_GetIntentPathClient, error)
}

type intentControllerClient struct {
	cc grpc.ClientConnInterface
}

func NewIntentControllerClient(cc grpc.ClientConnInterface) IntentControllerClient {
	return &intentControllerClient{cc}
}

func (c *intentControllerClient) GetIntentPath(ctx context.Context, opts ...grpc.CallOption) (IntentController_GetIntentPathClient, error) {
	stream, err := c.cc.NewStream(ctx, &IntentController_ServiceDesc.Streams[0], "/api.IntentController/GetIntentPath", opts...)
	if err != nil {
		return nil, err
	}
	x := &intentControllerGetIntentPathClient{stream}
	return x, nil
}

type IntentController_GetIntentPathClient interface {
	Send(*PathRequest) error
	Recv() (*PathResult, error)
	grpc.ClientStream
}

type intentControllerGetIntentPathClient struct {
	grpc.ClientStream
}

func (x *intentControllerGetIntentPathClient) Send(m *PathRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *intentControllerGetIntentPathClient) Recv() (*PathResult, error) {
	m := new(PathResult)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IntentControllerServer is the server API for IntentController service.
// All implementations must embed UnimplementedIntentControllerServer
// for forward compatibility
type IntentControllerServer interface {
	GetIntentPath(IntentController_GetIntentPathServer) error
	mustEmbedUnimplementedIntentControllerServer()
}

// UnimplementedIntentControllerServer must be embedded to have forward compatible implementations.
type UnimplementedIntentControllerServer struct {
}

func (UnimplementedIntentControllerServer) GetIntentPath(IntentController_GetIntentPathServer) error {
	return status.Errorf(codes.Unimplemented, "method GetIntentPath not implemented")
}
func (UnimplementedIntentControllerServer) mustEmbedUnimplementedIntentControllerServer() {}

// UnsafeIntentControllerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IntentControllerServer will
// result in compilation errors.
type UnsafeIntentControllerServer interface {
	mustEmbedUnimplementedIntentControllerServer()
}

func RegisterIntentControllerServer(s grpc.ServiceRegistrar, srv IntentControllerServer) {
	s.RegisterService(&IntentController_ServiceDesc, srv)
}

func _IntentController_GetIntentPath_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(IntentControllerServer).GetIntentPath(&intentControllerGetIntentPathServer{stream})
}

type IntentController_GetIntentPathServer interface {
	Send(*PathResult) error
	Recv() (*PathRequest, error)
	grpc.ServerStream
}

type intentControllerGetIntentPathServer struct {
	grpc.ServerStream
}

func (x *intentControllerGetIntentPathServer) Send(m *PathResult) error {
	return x.ServerStream.SendMsg(m)
}

func (x *intentControllerGetIntentPathServer) Recv() (*PathRequest, error) {
	m := new(PathRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IntentController_ServiceDesc is the grpc.ServiceDesc for IntentController service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IntentController_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.IntentController",
	HandlerType: (*IntentControllerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetIntentPath",
			Handler:       _IntentController_GetIntentPath_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/intent.proto",
}
