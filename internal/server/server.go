package server

import (
	"github.com/JueViGrace/bakery-server/internal/api"
	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	api api.Api
}

func New() *FiberServer {
	app := fiber.New(fiber.Config{
		ServerHeader: "BakeryServer",
		AppName:      "BakeryServer",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			res := types.RespondBadRequest(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		},
	})
	db := data.NewStorage()

	return &FiberServer{
		api: api.New(app, db),
	}
}

func (s *FiberServer) Init() (err error) {
	err = s.api.Init()
	if err != nil {
		return err
	}
	return
}
