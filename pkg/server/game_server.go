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

type GameServer struct {
	services *service.Service
}

func (g *GameServer) GetRoomsByCode(_ context.Context, request *genproto.GetRoomByCodeRequest) (*genproto.GetRoomByCodeResponse, error) {
	rooms, err := g.services.GetRoomByCodePattern(request.GetCode(), 1)
	if err != nil {
		logger.WarningLogger.Printf("can't get rooms by code pattern: %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't get rooms by code pattern: %v", err.Error())
	}

	var genprotoRooms []*genproto.Room
	for _, room := range rooms {
		genprotoRooms = append(genprotoRooms, conversions.ConvertRoom(room))
	}

	return &genproto.GetRoomByCodeResponse{Rooms: genprotoRooms}, nil
}

func (g *GameServer) PutInQueue(ctx context.Context, request *genproto.PutInQueueRequest) (*emptypb.Empty, error) {
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

	g.services.AddUser(userID, request.GetRoomId())
	return &emptypb.Empty{}, nil
}

func (g *GameServer) DeleteFromQueue(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
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

	g.services.Cancel(userID)

	return &emptypb.Empty{}, nil
}

func (g *GameServer) AddCall(ctx context.Context, request *genproto.AddCallRequest) (*emptypb.Empty, error) {
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

	game, err := g.services.AddCall(userID, request.GetOpponentId(), request.GetRoomId())
	if err != nil {
		logger.WarningLogger.Printf("can't get game info: %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't get game info: %v", err.Error())
	}

	if game.UserIdFirst == -1 {
		return &emptypb.Empty{}, nil
	}

	err = g.services.SendGame(game)

	if err != nil {
		logger.WarningLogger.Printf("can't send game info: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "can't send game info: %v", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (g *GameServer) DeleteCall(ctx context.Context, request *genproto.DeleteCallRequest) (*emptypb.Empty, error) {
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

	err = g.services.DeleteCall(userID, request.GetOpponentId())
	if err != nil {
		logger.WarningLogger.Println(err)
		return nil, status.Errorf(codes.NotFound, "opponent not found: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (g *GameServer) SendTime(ctx context.Context, request *genproto.SendTimeRequest) (*emptypb.Empty, error) {
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

	if err := g.services.UpdateTime(request.GetGameId(), userID, request.GetTime()); err != nil {
		logger.WarningLogger.Printf("can't update time: %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't update time: %v", err.Error())
	}

	go g.services.SendResult(request.GetGameId(), userID, request.GetTime())
	return &emptypb.Empty{}, nil
}

func (g *GameServer) StreamGame(_ *emptypb.Empty, server genproto.GameService_StreamGameServer) error {
	md, ok := metadata.FromIncomingContext(server.Context())
	if !ok {
		logger.WarningLogger.Println("can't get metadata")
		return status.Error(codes.Unauthenticated, "can't get metadata")
	}
	userIDS := md.Get(userID)
	if len(userIDS) != 1 {
		logger.WarningLogger.Println("can't get user-id from metadata")
		return status.Error(codes.Unauthenticated, "can't get user-id from metadata")
	}

	userID, err := strconv.ParseInt(userIDS[0], 10, 64)
	if err != nil {
		logger.WarningLogger.Printf("can't parse %s: %v", userIDS[0], err.Error())
		return status.Errorf(codes.Internal, "can't parse %s: %v", userIDS[0], err.Error())
	}

	userChannel := g.services.CreateUserChannel(userID)
	for resp := range userChannel {
		err = server.Send(resp)
		if err != nil {
			logger.WarningLogger.Printf("can't send game result: %v", err.Error())
			return status.Errorf(codes.NotFound, "can't send game result: %v", err.Error())
		}
	}

	return nil
}

func NewGameServer(services *service.Service) *GameServer {
	return &GameServer{services: services}
}
