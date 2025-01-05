package util

import (
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

func GetIdFromParams(idString string) (*uuid.UUID, error) {
	id, err := uuid.Parse(idString)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
