package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/middleware/middlewareHandler"
	"gitlab.com/chalermphanEur/modules/middleware/middlewareRepository"
	"gitlab.com/chalermphanEur/modules/middleware/middlewareUsecase"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type (
	server struct {
		app        *fiber.App
		db         *gorm.DB
		mongoDb    *mongo.Client
		cfg        *config.Config
		elastic    *elasticsearch.Client
		middleware middlewareHandler.IMiddleHandlerService
	}
)

func NewMiddleware(cfg *config.Config) middlewareHandler.IMiddleHandlerService {
	repo := middlewareRepository.NewMiddlewareRepository()
	usecase := middlewareUsecase.NewMiddlewareUsercase(repo)
	return middlewareHandler.NewMiddlewareHandlerService(cfg, usecase)
}

func (s *server) gracefulShutdown(ctx context.Context, quit <-chan os.Signal) {
	log.Printf("Start service : %s", s.cfg.App.Name)

	<-quit
	log.Printf("Shutdown service : %s", s.cfg.App.Name)

	_, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := s.app.Shutdown(); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
}

func (s *server) httpListening() {
	if err := s.app.Listen(s.cfg.App.Url); err != nil {
		log.Fatal("Server shutdown: ", err)
	}
}

// Now don't use mongo
func StartMongo(ptcx context.Context, cfg *config.Config, mongoDb *mongo.Client) {

	s := &server{
		app: fiber.New(
			fiber.Config{
				BodyLimit: 1049000,
			},
		),
		mongoDb:    mongoDb,
		cfg:        cfg,
		middleware: NewMiddleware(cfg),
	}

	
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	cfg.App.Name = "kbm"

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(ptcx, quit)

	s.httpListening()
}

func StartSqlServer(ptcx context.Context, cfg *config.Config, db *gorm.DB) {

	s := &server{
		app: fiber.New(
			fiber.Config{
				BodyLimit: 1049000,
			},
		),
		db:         db,
		cfg:        cfg,
		middleware: NewMiddleware(cfg),
	}

	

	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	switch cfg.App.Name {
	case "auth":
		s.authService()
	case "user":
		s.userService()
	
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go s.gracefulShutdown(ptcx, quit)

	s.httpListening()
}
