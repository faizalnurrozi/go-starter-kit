package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/faizalnurrozi/go-starter-kit/internal/cache"
	"github.com/faizalnurrozi/go-starter-kit/internal/config"
	"github.com/faizalnurrozi/go-starter-kit/internal/database"
	dto "github.com/faizalnurrozi/go-starter-kit/internal/dto/request"
	"github.com/faizalnurrozi/go-starter-kit/internal/grpc"
	"github.com/faizalnurrozi/go-starter-kit/internal/handler"
	"github.com/faizalnurrozi/go-starter-kit/internal/logger"
	"github.com/faizalnurrozi/go-starter-kit/internal/middleware"
	repository_impl "github.com/faizalnurrozi/go-starter-kit/internal/repository/impl"
	serviceimpl "github.com/faizalnurrozi/go-starter-kit/internal/service/impl"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize logger
	logger.Init(cfg.Log.Level)

	// Initialize database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize cache
	redis := cache.NewRedisClient(cfg)

	// Initialize gRPC server
	grpcServer := grpc.NewServer(cfg)
	go grpcServer.Start()

	// Initialize repositories
	userRepo := repository_impl.NewUserRepository(db)

	// Initialize services
	userService := serviceimpl.NewUserService(userRepo, redis)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	healthHandler := handler.NewHealthHandler()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
	})

	// Global middleware
	app.Use(cors.New())
	app.Use(middleware.Logger())

	// Setup routes
	setupRoutes(app, userHandler, healthHandler)

	// Start server
	go func() {
		if err := app.Listen(":" + cfg.Server.Port); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	grpcServer.Stop()
	database.Close(db)
	redis.Close()

	logger.Info("Server exited")
}

func setupRoutes(app *fiber.App, userHandler *handler.UserHandler, healthHandler *handler.HealthHandler) {
	// Health check
	app.Get("/health", healthHandler.Check)

	// API versioning
	api := app.Group("/api")

	// V1 Routes
	v1 := api.Group("/v1")

	// User routes
	users := v1.Group("/users")
	users.Use(middleware.Auth()) // Auth middleware
	users.Get("/", userHandler.GetAll)
	users.Post("/", middleware.ValidateRequest(&dto.CreateUserRequest{}), userHandler.Create)
	users.Get("/:id", middleware.ValidateParams(), userHandler.GetByID)
	users.Put("/:id", middleware.ValidateParams(), middleware.ValidateRequest(&dto.UpdateUserRequest{}), userHandler.Update)
	users.Delete("/:id", middleware.ValidateParams(), userHandler.Delete)

	// V2 Routes (for future versions)
	v2 := api.Group("/v2")
	v2.Get("/users", userHandler.GetAll) // Same handler, different version
}
