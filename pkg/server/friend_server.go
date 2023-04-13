package server

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/conversions"
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
)

type FriendServer struct {
	services *service.Service
}

func (f *FriendServer) AddFriend(ctx context.Context, request *genproto.AddFriendRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.WarningLogger.Println("can't get metadata")
		return nil, status.Error(codes.Unauthenticated, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		logger.WarningLogger.Println("can't get user-id from metadata")
		return nil, status.Error(codes.Unauthenticated, "can't get user-id from metadata")
	}

	userID, err := strconv.ParseInt(userIDS[0], 10, 64)
	if err != nil {
		logger.WarningLogger.Printf("can't parse %s: %v", userIDS[0], err.Error())
		return nil, status.Errorf(codes.Internal, "can't parse %s: %v", userIDS[0], err.Error())
	}

	err = f.services.AddFriend(userID, request.GetUserId())
	if err != nil {
		logger.WarningLogger.Printf("can't add friend %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't add friend %v", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (f *FriendServer) DeleteFriend(ctx context.Context, request *genproto.DeleteFriendRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.WarningLogger.Println("can't get metadata")
		return nil, status.Error(codes.Unauthenticated, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		logger.WarningLogger.Println("can't get user-id from metadata")
		return nil, status.Error(codes.Unauthenticated, "can't get user-id from metadata")
	}

	userID, err := strconv.ParseInt(userIDS[0], 10, 64)
	if err != nil {
		logger.WarningLogger.Printf("can't parse %s: %v", userIDS[0], err.Error())
		return nil, status.Errorf(codes.Internal, "can't parse %s: %v", userIDS[0], err.Error())
	}

	err = f.services.DeleteFriend(userID, request.GetUserId())
	if err != nil {
		logger.WarningLogger.Printf("can't add friend %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't add friend %v", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (f *FriendServer) GetFriends(ctx context.Context, _ *emptypb.Empty) (*genproto.Users, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.WarningLogger.Println("can't get metadata")
		return nil, status.Error(codes.Unauthenticated, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		logger.WarningLogger.Println("can't get user-id from metadata")
		return nil, status.Error(codes.Unauthenticated, "can't get user-id from metadata")
	}

	userID, err := strconv.ParseInt(userIDS[0], 10, 64)
	if err != nil {
		logger.WarningLogger.Printf("can't parse %s: %v", userIDS[0], err.Error())
		return nil, status.Errorf(codes.Internal, "can't parse %s: %v", userIDS[0], err.Error())
	}

	users, err := f.services.GetFriends(userID)
	if err != nil {
		logger.WarningLogger.Printf("can't find users: %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't find users: %v", err.Error())
	}

	var protoUsers []*genproto.User
	for _, user := range users {
		protoUsers = append(protoUsers, conversions.ConvertUser(user))
	}
	return &genproto.Users{Users: protoUsers}, err
}

func NewFriendServer(services *service.Service) *FriendServer {
	return &FriendServer{services: services}
}
