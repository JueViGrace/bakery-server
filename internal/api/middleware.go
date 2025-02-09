package api

import (
	"github.com/JueViGrace/bakery-server/internal/data"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/gofiber/fiber/v2"
)

func (a *api) adminAuthMiddleware(c *fiber.Ctx) error {
	data, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	if data.Role != types.Admin {
		res := types.RespondForbbiden(nil, "forbbiden resource")
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) sessionMiddleware(c *fiber.Ctx) error {
	_, err := getUserDataForReq(c, a.db)
	if err != nil {
		res := types.RespondUnauthorized(nil, err.Error())
		return c.Status(res.Status).JSON(res)
	}

	return c.Next()
}

func (a *api) authenticatedHandler(handler types.AuthDataHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		data, err := getUserDataForReq(c, a.db)
		if err != nil {
			res := types.RespondUnauthorized(nil, err.Error())
			return c.Status(res.Status).JSON(res)
		}

		return handler(c, data)
	}
}

func getUserDataForReq(c *fiber.Ctx, db data.Storage) (*types.AuthData, error) {
	jwt, err := util.ExtractJWTFromHeader(c, func(s string) {
		db.SessionStore().DeleteSessionByToken(s)
	})
	if err != nil {
		return nil, err
	}

	session, err := db.SessionStore().GetSessionById(jwt.Claims.SessionID)
	if err != nil {
		return nil, err
	}

	dbUser, err := db.UserStore().GetUserById(&session.UserId)
	if err != nil {
		return nil, err
	}

	return &types.AuthData{
		UserId:    dbUser.ID,
		SessionId: session.ID,
		Username:  dbUser.Username,
		Role:      dbUser.Role,
	}, nil
}
