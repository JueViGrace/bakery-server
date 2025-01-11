package api

import (
	"github.com/JueViGrace/bakery-server/internal/handlers"
	"github.com/JueViGrace/bakery-server/internal/types"
	"github.com/JueViGrace/bakery-server/internal/util"
	"github.com/gofiber/fiber/v2"
)

func (a *api) UserRoutes(api fiber.Router) {
	usersGroup := api.Group("/users", a.sessionMiddleware)

	userHandler := handlers.NewUserHandler(a.db.UserStore())

	usersGroup.Get("/", a.sessionMiddleware, a.adminAuthMiddleware, userHandler.GetUsers)
	usersGroup.Get("/me", a.sessionMiddleware, func(c *fiber.Ctx) error {
		jwt, err := util.ExtractJWTFromHeader(c, func(s string) {
			a.db.SessionStore().DeleteTokenByToken(s)
		})
		if err != nil {
			res := types.RespondBadRequest(err.Error(), "invalid request")
			return c.Status(res.Status).JSON(res)
		}
		return userHandler.GetUserById(c, &jwt.Claims.UserId)
	})
	usersGroup.Patch("/:id", a.sessionMiddleware, a.userIdMiddleware, userHandler.UpdateUser)
	usersGroup.Delete("/:id", a.sessionMiddleware, a.userIdMiddleware, userHandler.DeleteUser)
}
