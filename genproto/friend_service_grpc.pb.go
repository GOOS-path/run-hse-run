// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: friend_service.proto

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

// FriendServiceClient is the client API for FriendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FriendServiceClient interface {
	AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Users, error)
}

type friendServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendServiceClient(cc grpc.ClientConnInterface) FriendServiceClient {
	return &friendServiceClient{cc}
}

func (c *friendServiceClient) AddFriend(ctx context.Context, in *AddFriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.FriendService/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) DeleteFriend(ctx context.Context, in *DeleteFriendRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/run.hse.run.FriendService/DeleteFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *friendServiceClient) GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/run.hse.run.FriendService/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServiceServer is the server API for FriendService service.
// All implementations should embed UnimplementedFriendServiceServer
// for forward compatibility
type FriendServiceServer interface {
	AddFriend(context.Context, *AddFriendRequest) (*emptypb.Empty, error)
	DeleteFriend(context.Context, *DeleteFriendRequest) (*emptypb.Empty, error)
	GetFriends(context.Context, *emptypb.Empty) (*Users, error)
}

// UnimplementedFriendServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFriendServiceServer struct {
}

func (UnimplementedFriendServiceServer) AddFriend(context.Context, *AddFriendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedFriendServiceServer) DeleteFriend(context.Context, *DeleteFriendRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFriend not implemented")
}
func (UnimplementedFriendServiceServer) GetFriends(context.Context, *emptypb.Empty) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}

// UnsafeFriendServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FriendServiceServer will
// result in compilation errors.
type UnsafeFriendServiceServer interface {
	mustEmbedUnimplementedFriendServiceServer()
}

func RegisterFriendServiceServer(s grpc.ServiceRegistrar, srv FriendServiceServer) {
	s.RegisterService(&FriendService_ServiceDesc, srv)
}

func _FriendService_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.FriendService/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).AddFriend(ctx, req.(*AddFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_DeleteFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFriendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).DeleteFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.FriendService/DeleteFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).DeleteFriend(ctx, req.(*DeleteFriendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FriendService_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServiceServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/run.hse.run.FriendService/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServiceServer).GetFriends(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// FriendService_ServiceDesc is the grpc.ServiceDesc for FriendService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FriendService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "run.hse.run.FriendService",
	HandlerType: (*FriendServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddFriend",
			Handler:    _FriendService_AddFriend_Handler,
		},
		{
			MethodName: "DeleteFriend",
			Handler:    _FriendService_DeleteFriend_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _FriendService_GetFriends_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend_service.proto",
}
