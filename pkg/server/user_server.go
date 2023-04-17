package server

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/conversions"
	"Run_Hse_Run/pkg/logger"
	"Run_Hse_Run/pkg/service"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"os"
	"sort"
	"strconv"
)

type UserServer struct {
	services *service.Service
}

func createFile(userId int64) {
	file, err := os.Create(fmt.Sprintf("images/image_%d", userId))
	if err != nil {
		logger.WarningLogger.Println(err.Error())
		return
	}
	file.Close()
}

func saveImage(userId int64, content string) {
	file, err := os.OpenFile(fmt.Sprintf("images/image_%d", userId), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.WarningLogger.Println(err.Error())
		return
	}
	defer file.Close()

	_, _ = file.WriteString(content)
}

func getImage(userId int64) string {
	file, err := os.OpenFile(fmt.Sprintf("images/image_%d", userId), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logger.WarningLogger.Println(err.Error())
		return ""
	}
	defer file.Close()

	output, _ := io.ReadAll(file)
	return string(output)
}

func (u *UserServer) GetUserByID(_ context.Context, request *genproto.GetUserByIDRequest) (*genproto.User, error) {
	user, err := u.services.GetUserById(request.GetId())
	if err != nil {
		logger.WarningLogger.Println(err)
		return nil, status.Errorf(codes.NotFound, "user with id: %d not found: %v", request.GetId(), err.Error())
	}

	user.Image = getImage(user.Id)
	return conversions.ConvertUser(user), nil
}

func (u *UserServer) GetLeaderBoard(_ context.Context, _ *emptypb.Empty) (*genproto.Users, error) {
	users, err := u.services.GetUsers()
	if err != nil {
		logger.WarningLogger.Printf("can't get users: %v", err.Error())
		return nil, status.Errorf(codes.Internal, "can't get users: %v", err.Error())
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Score > users[j].Score
	})

	var protoUsers []*genproto.User
	for _, user := range users {
		user.Image = getImage(user.Id)
		protoUsers = append(protoUsers, conversions.ConvertUser(user))
	}
	return &genproto.Users{Users: protoUsers}, err
}

func (u *UserServer) GetMe(ctx context.Context, _ *emptypb.Empty) (*genproto.User, error) {
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

	user, err := u.services.GetUserById(userID)
	if err != nil {
		logger.WarningLogger.Println(err)
		return nil, status.Errorf(codes.NotFound, "user with id: %d not found: %v", userID, err.Error())
	}

	user.Image = getImage(user.Id)

	return conversions.ConvertUser(user), nil
}

func (u *UserServer) GetUserByNickname(_ context.Context, request *genproto.GetUserByNicknameRequest) (*genproto.Users, error) {
	users, err := u.services.GetUsersByNicknamePattern(request.GetNickname())
	if err != nil {
		logger.WarningLogger.Printf("can't get user by pattern: %v", err.Error())
		return nil, status.Errorf(codes.NotFound, "can't get user by pattern: %v", err.Error())
	}

	var protoUsers []*genproto.User
	for _, user := range users {
		user.Image = getImage(user.Id)
		protoUsers = append(protoUsers, conversions.ConvertUser(user))
	}
	return &genproto.Users{Users: protoUsers}, err
}

func (u *UserServer) ChangeNickname(ctx context.Context, request *genproto.ChangeNicknameRequest) (*emptypb.Empty, error) {
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

	err = u.services.RenameUser(userID, request.GetNewNickname())
	if err != nil {
		logger.WarningLogger.Printf("can't rename user: %v", err.Error())
		if err.Error() == service.NicknameError {
			return nil, status.Error(codes.InvalidArgument, "invalid nickname")
		} else {
			return nil, status.Errorf(codes.Internal, "can't rename user: %v", err.Error())
		}
	}

	return &emptypb.Empty{}, nil
}

func (u *UserServer) ChangeImage(ctx context.Context, request *genproto.ChangeImageRequest) (*emptypb.Empty, error) {
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

	saveImage(userID, request.GetNewImage())

	return &emptypb.Empty{}, nil
}

func NewUserServer(services *service.Service) *UserServer {
	return &UserServer{services: services}
}
