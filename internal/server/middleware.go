package server

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/labstack/echo/v4"
)

const bypassUserID = "clwi6195d0000w6bqu9vhnc0h"

func sessionCheckMiddleware(client *db.PrismaClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(_c echo.Context) error {
			c := _c.(*handlers.Context)

			user, err := client.User.FindFirst(db.User.ID.Equals(bypassUserID)).Exec(c.Request().Context())
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			c.SetUser(user)

			return next(c)
		}
	}
}
