package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberServer struct {
	*fiber.App

	db data.Service
}

func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "BakeryServer",
			AppName:      "BakeryServer",
		}),

		db: data.NewService(),
	}
}

func (s *FiberServer) Init() (err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return err
	}

	s.RegisterRoutes()
	s.App.Use(cors.New())
	s.App.Use(logger.New())

	err = s.Listen(fmt.Sprintf(":%d", port))

	return
}
