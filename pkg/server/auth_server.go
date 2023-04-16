package server

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/conversions"
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/model"
	"Run_Hse_Run/pkg/service"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthServer struct {
	services *service.Service
}

func (a *AuthServer) Registration(_ context.Context, request *genproto.User) (*genproto.UserWithToken, error) {
	_, err := a.services.GetUser(request.GetEmail())
	if err == nil {
		logger.WarningLogger.Println("user already exist")
		return nil, status.Errorf(codes.AlreadyExists, "user already exist with email: %s", request.GetEmail())
	}

	user := model.User{
		Email:    request.GetEmail(),
		Nickname: request.GetNickname(),
		Image:    request.GetImage(),
		Score:    0,
	}

	id, err := a.services.CreateUser(user)

	if err != nil {
		logger.WarningLogger.Println(err)
		if err.Error() == service.NicknameError {
			return nil, status.Error(codes.InvalidArgument, "invalid nickname")
		}
		return nil, status.Errorf(codes.Internal, "server internal error: %v", err.Error())
	}

	user.Id = id

	token, err := a.services.GenerateToken(request.GetEmail())
	if err != nil {
		logger.WarningLogger.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "invalid nickname: %v", err.Error())
	}

	return &genproto.UserWithToken{
		AccessToken: token,
		User:        conversions.ConvertUser(user),
	}, nil
}

func (a *AuthServer) SendVerifyEmail(_ context.Context, request *genproto.SendVerifyEmailRequest) (*emptypb.Empty, error) {
	if request.GetEmail() == "" {
		logger.WarningLogger.Println("invalid email")
		return nil, status.Error(codes.InvalidArgument, "invalid email")
	}

	if err := a.services.SendEmail(request.GetEmail()); err != nil {
		logger.WarningLogger.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "can't send verify email: %v", err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (a *AuthServer) Verify(_ context.Context, request *genproto.VerifyRequest) (*genproto.VerifyResponse, error) {
	if request.GetCode() == 0 || request.GetEmail() == "" {
		logger.WarningLogger.Println("invalid email or code")
		return nil, status.Error(codes.Internal, "invalid email or code")
	}

	service.Mu.Lock()
	code, ok := service.Codes[request.GetEmail()]
	service.Mu.Unlock()

	if !ok {
		logger.WarningLogger.Println("email didn't added")
		return nil, status.Error(codes.NotFound, "email doesn't exist")
	}

	if code != request.GetCode() {
		logger.WarningLogger.Println("incorrect code")
		return nil, status.Error(codes.NotFound, "code doesn't much")
	}

	user, err := a.services.GetUser(request.GetEmail())
	if err != nil {
		logger.WarningLogger.Println("user doesn't exist in db")
		return &genproto.VerifyResponse{
			Response: &genproto.VerifyResponse_NeedRegistration{
				NeedRegistration: true,
			},
		}, nil
	}

	token, err := a.services.GenerateToken(request.GetEmail())
	if err != nil {
		logger.WarningLogger.Println(err)
		return nil, status.Errorf(codes.InvalidArgument, "can't generate token: %v", err.Error())
	}

	return &genproto.VerifyResponse{
		Response: &genproto.VerifyResponse_UserInfo{
			UserInfo: &genproto.UserWithToken{
				AccessToken: token,
				User:        conversions.ConvertUser(user),
			},
		},
	}, nil
}

func NewAuthServer(services *service.Service) *AuthServer {
	return &AuthServer{services: services}
}
