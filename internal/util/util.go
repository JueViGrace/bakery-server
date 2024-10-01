package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetIdFromParams(ctx *fiber.Ctx) (*uuid.UUID, error) {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func HashPassword(password string) (*string, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

    pass := string(encpw)

	return &pass, nil
}

func ValidatePassword(reqPass string, encPass string) bool {
    return bcrypt.CompareHashAndPassword([]byte(encPass), []byte(reqPass)) == nil
}
