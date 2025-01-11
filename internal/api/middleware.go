package api

import (
	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/gofiber/fiber/v2"
)

type AuthData struct {
	jwt  util.JwtData
	role string
}

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.role != types.Admin {
		res := types.RespondForbbiden("permission denied", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	_, err = a.db.SessionStore().GetTokenByToken(data.jwt.Token.Raw)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) userIdMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.role == types.Admin {
		return c.Next()
	}

	id, err := util.GetIdFromParams(c.Params("id"))
	if err != nil {
		res := types.RespondBadRequest(err.Error(), "Failed")
		return c.Status(res.Status).JSON(res)
	}

	if data.jwt.Claims.UserId != *id {
		res := types.RespondForbbiden("Forbbiden action", "Failed")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

// // todo: orders middleware

func getUserDataForReq(c *fiber.Ctx, db data.Storage) (*AuthData, error) {
	jwt, err := util.ExtractJWTFromHeader(c, func(s string) {
		db.SessionStore().DeleteTokenByToken(s)
	})
	if err != nil {
		return nil, err
	}

	user, err := db.UserStore().GetUserById(&jwt.Claims.UserId)
	if err != nil {
		return nil, err
	}

	return &AuthData{
		jwt:  *jwt,
		role: user.Role,
	}, nil
}
