package grpc

import (
	"net"

	"github.com/faizalnurrozi/go-starter-kit/internal/config"
	"github.com/faizalnurrozi/go-starter-kit/internal/grpc/handlers"
	"github.com/faizalnurrozi/go-starter-kit/internal/logger"
	repository_impl "github.com/faizalnurrozi/go-starter-kit/internal/repository/impl"
	serviceimpl "github.com/faizalnurrozi/go-starter-kit/internal/service/impl"
	pb "github.com/faizalnurrozi/go-starter-kit/proto/user"
	"github.com/faizalnurrozi/go-starter-kit/internal/database"

	"google.golang.org/grpc"
)

type Server struct {
	server *grpc.Server
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	grpcServer := grpc.NewServer()

	// Initialize database
	db, err := database.Connect(cfg)
	if err != nil {
		logger.Fatal("Failed to connect database:", err)
	}

	// Inisialisasi dependencies
	userRepo := repository_impl.NewUserRepository(db)
	userService := serviceimpl.NewUserService(userRepo, nil)
	userHandler := handlers.NewUserHandler(userService)

	// Registrasi handler
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	return &Server{
		server: grpcServer,
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
