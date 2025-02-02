package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Api interface {
	Init() error
}

type api struct {
	*fiber.App
	db        data.Storage
	validator *util.XValidator
}

func New() Api {
	return &api{
		App: fiber.New(fiber.Config{
			ServerHeader: "BakeryServer",
			AppName:      "BakeryServer",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				res := types.RespondBadRequest(nil, err.Error())
				return c.Status(res.Status).JSON(res)
			},
		}),
		db: data.NewStorage(),
		validator: &util.XValidator{
			Validator: validator.New(),
		},
	}
}

func (a *api) Init() (err error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	a.App.Use(logger.New())
	a.App.Use(cors.New())

	a.RegisterRoutes()

	a.App.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	err = a.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		a.db.Close()
		return err
	}

	return
}
