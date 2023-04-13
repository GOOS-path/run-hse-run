// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: game_server.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	GetRoomsByCode(ctx context.Context, in *GetRoomByCodeRequest, opts ...grpc.CallOption) (*GetRoomByCodeResponse, error)
	PutInQueue(ctx context.Context, in *PutInQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteFromQueue(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddCall(ctx context.Context, in *AddCallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteCall(ctx context.Context, in *DeleteCallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SendTime(ctx context.Context, in *SendTimeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StreamGame(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (GameService_StreamGameClient, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) GetRoomsByCode(ctx context.Context, in *GetRoomByCodeRequest, opts ...grpc.CallOption) (*GetRoomByCodeResponse, error) {
	out := new(GetRoomByCodeResponse)
	err := c.cc.Invoke(ctx, "/run.hse.run.GameService/GetRoomsByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) PutInQueue(ctx context.Context, in *PutInQueueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.GameService/PutInQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) DeleteFromQueue(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.GameService/DeleteFromQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) AddCall(ctx context.Context, in *AddCallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.GameService/AddCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) DeleteCall(ctx context.Context, in *DeleteCallRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.GameService/DeleteCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) SendTime(ctx context.Context, in *SendTimeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.GameService/SendTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) StreamGame(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (GameService_StreamGameClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[0], "/run.hse.run.GameService/StreamGame", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceStreamGameClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_StreamGameClient interface {
	Recv() (*StreamResponse, error)
	grpc.ClientStream
}

type gameServiceStreamGameClient struct {
	grpc.ClientStream
}

func (x *gameServiceStreamGameClient) Recv() (*StreamResponse, error) {
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations should embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	GetRoomsByCode(context.Context, *GetRoomByCodeRequest) (*GetRoomByCodeResponse, error)
	PutInQueue(context.Context, *PutInQueueRequest) (*emptypb.Empty, error)
	DeleteFromQueue(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	AddCall(context.Context, *AddCallRequest) (*emptypb.Empty, error)
	DeleteCall(context.Context, *DeleteCallRequest) (*emptypb.Empty, error)
	SendTime(context.Context, *SendTimeRequest) (*emptypb.Empty, error)
	StreamGame(*emptypb.Empty, GameService_StreamGameServer) error
}

// UnimplementedGameServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) GetRoomsByCode(context.Context, *GetRoomByCodeRequest) (*GetRoomByCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoomsByCode not implemented")
}
func (UnimplementedGameServiceServer) PutInQueue(context.Context, *PutInQueueRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutInQueue not implemented")
}
func (UnimplementedGameServiceServer) DeleteFromQueue(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFromQueue not implemented")
}
func (UnimplementedGameServiceServer) AddCall(context.Context, *AddCallRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCall not implemented")
}
func (UnimplementedGameServiceServer) DeleteCall(context.Context, *DeleteCallRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCall not implemented")
}
func (UnimplementedGameServiceServer) SendTime(context.Context, *SendTimeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTime not implemented")
}
func (UnimplementedGameServiceServer) StreamGame(*emptypb.Empty, GameService_StreamGameServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamGame not implemented")
}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_GetRoomsByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).GetRoomsByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.GameService/GetRoomsByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).GetRoomsByCode(ctx, req.(*GetRoomByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_PutInQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutInQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).PutInQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.GameService/PutInQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).PutInQueue(ctx, req.(*PutInQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_DeleteFromQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).DeleteFromQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.GameService/DeleteFromQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).DeleteFromQueue(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_AddCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).AddCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.GameService/AddCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).AddCall(ctx, req.(*AddCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_DeleteCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).DeleteCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.GameService/DeleteCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).DeleteCall(ctx, req.(*DeleteCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_SendTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).SendTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.GameService/SendTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).SendTime(ctx, req.(*SendTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_StreamGame_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).StreamGame(m, &gameServiceStreamGameServer{stream})
}

type GameService_StreamGameServer interface {
	Send(*StreamResponse) error
	grpc.ServerStream
}

type gameServiceStreamGameServer struct {
	grpc.ServerStream
}

func (x *gameServiceStreamGameServer) Send(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "run.hse.run.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRoomsByCode",
			Handler:    _GameService_GetRoomsByCode_Handler,
		},
		{
			MethodName: "PutInQueue",
			Handler:    _GameService_PutInQueue_Handler,
		},
		{
			MethodName: "DeleteFromQueue",
			Handler:    _GameService_DeleteFromQueue_Handler,
		},
		{
			MethodName: "AddCall",
			Handler:    _GameService_AddCall_Handler,
		},
		{
			MethodName: "DeleteCall",
			Handler:    _GameService_DeleteCall_Handler,
		},
		{
			MethodName: "SendTime",
			Handler:    _GameService_SendTime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamGame",
			Handler:       _GameService_StreamGame_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "game_server.proto",
}