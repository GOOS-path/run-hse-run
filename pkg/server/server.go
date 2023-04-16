package server

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPCServer struct {
	grpcServer *grpc.Server
	services   *service.Service
}

func NewGRPCServer(services *service.Service) *GRPCServer {
	return &GRPCServer{services: services}
}

func (srv *GRPCServer) Run(listener net.Listener) error {
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(srv.unaryInterceptor()),
	}
	srv.grpcServer = grpc.NewServer(serverOptions...)

	authServer := NewAuthServer(srv.services)
	userServer := NewUserServer(srv.services)
	friendServer := NewFriendServer(srv.services)
	gameServer := NewGameServer(srv.services)

	genproto.RegisterAuthServiceServer(srv.grpcServer, authServer)
	genproto.RegisterUserServiceServer(srv.grpcServer, userServer)
	genproto.RegisterFriendServiceServer(srv.grpcServer, friendServer)
	genproto.RegisterGameServiceServer(srv.grpcServer, gameServer)
	reflection.Register(srv.grpcServer)

	return srv.grpcServer.Serve(listener)
}
