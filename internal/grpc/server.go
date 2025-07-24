package grpc

import (
	"net"

	"github.com/faizalnurrozi/go-starter-kit/internal/config"
	"github.com/faizalnurrozi/go-starter-kit/internal/logger"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	server := grpc.NewServer()

	// Register your gRPC services here
	// pb.RegisterUserServiceServer(server, &userGRPCHandler{})

	return &Server{
		server: server,
		config: cfg,
	}
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp", ":"+s.config.GRPC.Port)
	if err != nil {
		logger.Fatal("Failed to listen for gRPC:", err)
	}

	logger.Info("gRPC server listening on port " + s.config.GRPC.Port)

	if err := s.server.Serve(lis); err != nil {
		logger.Fatal("Failed to serve gRPC:", err)
	}
}

func (s *Server) Stop() {
	s.server.GracefulStop()
}
