package server

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/labstack/echo/v4"
)

func sessionCheckMiddleware(client *db.PrismaClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(_c echo.Context) error {
			c := _c.(*handlers.Context)

			user, err := client.User.FindFirst().Exec(c.Request().Context())
			if db.IsErrNotFound(err) {
				user, err = client.User.CreateOne(db.User.Name.Set("test user")).Exec(c.Request().Context())
			}
			if err != nil {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			c.SetUser(user)

			return next(c)
		}
	}
}
