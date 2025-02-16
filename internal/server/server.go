package server

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/JueViGrace/bakery-server/internal/api"
	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberServer struct {
	*fiber.App
	db  data.Storage
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
		App: app,
		api: api.New(app, db),
	}
}

func (s *FiberServer) Init() (err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	s.App.Use(logger.New())
	s.App.Use(cors.New())
	s.Static("/static", os.Getenv("STATIC_RES"), fiber.Static{
		Compress:      true,
		CacheDuration: 10 * time.Minute,
	})

	err = s.api.Init()
	if err != nil {
		return err
	}

	s.App.Use(func(c *fiber.Ctx) error {
		res := types.RespondNotFound(nil, "Not found")
		return c.Status(res.Status).JSON(res)
	})

	err = s.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		s.db.Close()
		return err
	}
	return
}
