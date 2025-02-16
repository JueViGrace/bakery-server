package api

import (
	"fmt"
	"os"
	"strconv"
	"time"

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

func New(app *fiber.App, db data.Storage) Api {
	return &api{
		App: app,
		db:  db,
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

	a.Static("/static", os.Getenv("STATIC_RES"), fiber.Static{
		Compress:      true,
		CacheDuration: 10 * time.Minute,
	})

	a.RegisterRoutes()

	a.App.Use(func(c *fiber.Ctx) error {
		res := types.RespondNotFound(nil, "Not found")
		return c.Status(res.Status).JSON(res)
	})

	err = a.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		a.db.Close()
		return err
	}

	return
}
