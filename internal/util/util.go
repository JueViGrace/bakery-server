package util

import (
	"fmt"
	"time"

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

func FormatDateForResponse(d time.Time) string {
	return fmt.Sprintf(
		"%d-%02d-%02d %02d:%02d:%02d",
		d.Year(), d.Month(), d.Day(),
		d.Hour(), d.Minute(), d.Second(),
	)
}
