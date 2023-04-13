package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationKey = "authorization"
	userID           = "user-id"
)

var accessibleRoles = map[string]struct{}{
	"/run.hse.run.AuthService/Registration":    {},
	"/run.hse.run.AuthService/SendVerifyEmail": {},
	"/run.hse.run.AuthService/Verify":          {},
}

func (srv *GRPCServer) unaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		_, ok := accessibleRoles[info.FullMethod]
		if ok {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		values := md[authorizationKey]
		if len(values) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
		}

		accessToken := values[0]
		userID, err := srv.services.ParseToken(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
		}

		md.Append("user-id", fmt.Sprintf("%d", userID))
		ctx = metadata.NewIncomingContext(ctx, md)

		return handler(ctx, req)
	}
}