package api

import (
	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	a.RegisterRoutes()

	return
}
