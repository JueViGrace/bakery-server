package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetIdFromParams(ctx *fiber.Ctx) (*uuid.UUID, error) {
    id, err := uuid.Parse(ctx.Params("id"))
    if err != nil {
        return nil, err
    }

    return &id, nil
}
