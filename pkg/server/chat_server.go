package server

import (
	"Run_Hse_Run/genproto"
	"Run_Hse_Run/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"strconv"
	"sync"
)

type ChatServer struct {
	mu   sync.Mutex
	chat map[int64]chan *genproto.MessageResponse
}

func (c *ChatServer) DoChatting(server genproto.ChatService_DoChattingServer) error {
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

	c.mu.Lock()
	if _, ok := c.chat[userID]; !ok {
		c.chat[userID] = make(chan *genproto.MessageResponse)
	}
	c.mu.Unlock()
	defer func() {
		c.mu.Lock()
		close(c.chat[userID])
		c.mu.Unlock()
	}()
	defer func() {
		c.mu.Lock()
		delete(c.chat, userID)
		c.mu.Unlock()
	}()

	go func() {
		for message := range c.chat[userID] {
			if err := server.Send(message); err != nil {
				logger.WarningLogger.Printf("can't send message to user: %v", err.Error())
			}
		}
	}()

	for {
		in, err := server.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return status.Errorf(codes.Internal, "can't receive message: %v", err.Error())
		}

		c.mu.Lock()
		if ch, ok := c.chat[in.GetUserTo()]; !ok {
			c.mu.Unlock()
			return status.Error(codes.NotFound, "user isn't in application")
		} else {
			ch <- &genproto.MessageResponse{
				Content:  in.GetContent(),
				UserFrom: userID,
			}
			c.mu.Unlock()
		}
	}
}

func NewChatServer() *ChatServer {
	return &ChatServer{
		chat: make(map[int64]chan *genproto.MessageResponse),
	}
}
