package server

import (
	"context"
	"net/http"
	"time"

	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func sessionCheckMiddleware(client *db.PrismaClient) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(_c echo.Context) error {
			c := _c.(*handlers.Context)

			sessionCookie, err := c.Cookie(logic.SessionCookieName)
			if err != nil || sessionCookie.Value == "" {
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}

			session, err := logic.GetAndVerifyUserSession(c.Request().Context(), client, sessionCookie.Value, logic.StringToIdentifier(c.RealIP()))
			if err != nil {
				if !db.IsErrNotFound(err) {
					log.Err(err).Str("sessionValue", sessionCookie.Value).Msg("failed to retrieve a session")
				}

				blank := logic.NewSessionCookie("", time.Unix(0, 0))
				c.SetCookie(&blank)

				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}

			if time.Until(session.Expiration) < time.Hour*24 {
				// Extend session by another 7 days
				go func() {
					ctx, cancel := context.WithTimeout(context.Background(), time.Second)
					defer cancel()
					_, err := logic.UpdateSessionExpiration(ctx, client, session.ID, logic.SessionExpiration7Days())
					if err != nil {
						log.Err(err).Str("sessionId", session.ID).Msg("failed to update session expiration")
					}
				}()
			}

			c.SetUser(session.User())
			return next(c)
		}
	}
}
