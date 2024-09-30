package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JueViGrace/bakery-go/internal/data"
	"github.com/gofiber/fiber/v2"
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
        
	err = s.Listen(fmt.Sprintf(":%d", port))

	return
}
