package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JueViGrace/bakery-go/internal/database"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	return &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "BakeryServer",
			AppName:      "BakeryServer",
		}),

		db: database.New(),
	}
}

func (s *FiberServer) Init() (err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return err
	}

	err = s.Listen(fmt.Sprintf(":%d", port))

	return
}
