package server

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

type APIResponse struct {
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Data        any       `json:"data"`
	Time        time.Time `json:"time"`
	Message     string    `json:"message"`
}

func NewAPIResponse(status int, description string, data any, message string) *APIResponse {
	return &APIResponse{
		Status:      status,
		Description: description,
		Data:        data,
		Time:        time.Now().UTC(),
		Message:     message,
	}
}

func RespondOk(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusOK).JSON(NewAPIResponse(http.StatusOK, http.StatusText(http.StatusOK), data, message))
}

func RespondCreated(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusCreated).JSON(NewAPIResponse(http.StatusCreated, http.StatusText(http.StatusCreated), data, message))
}

func RespondAccepted(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusAccepted).JSON(NewAPIResponse(http.StatusAccepted, http.StatusText(http.StatusAccepted), data, message))
}

func RespondNoContent(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusNoContent).JSON(NewAPIResponse(http.StatusNoContent, http.StatusText(http.StatusNoContent), data, message))
}

func RespondBadRequest(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusBadRequest).JSON(NewAPIResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), data, message))
}

func RespondUnauthorized(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusUnauthorized).JSON(NewAPIResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), data, message))
}

func RespondForbbiden(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusForbidden).JSON(NewAPIResponse(http.StatusForbidden, http.StatusText(http.StatusForbidden), data, message))
}

func RespondNotFound(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusNotFound).JSON(NewAPIResponse(http.StatusNotFound, http.StatusText(http.StatusForbidden), data, message))
}

func RespondInternalServerError(c *fiber.Ctx, data any, message string) error {
	return c.Status(http.StatusInternalServerError).JSON(NewAPIResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), data, message))
}
