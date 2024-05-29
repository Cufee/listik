package server

import (
	"io/fs"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/rs/zerolog"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/server/handlers/api"
	"github.com/cufee/shopping-list/internal/server/handlers/app"
	"github.com/cufee/shopping-list/internal/templates/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Create a new echo.Echo instance with all routes registered
func New(db *db.PrismaClient, assets fs.FS) *echo.Echo {
	e := echo.New()
	if assets != nil {
		e.StaticFS("/static", assets)
	}

	// echo does not route correctly when a `/` route is attached to a group, this is to fix that issue
	e.Pre(middleware.AddTrailingSlash())

	logger := zerolog.New(os.Stdout)
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogError:   true,
		LogStatus:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Int("status", v.Status).
				Int64("duration_ms", v.Latency.Milliseconds()).
				Str("URI", v.URI).
				Err(v.Error).
				Msg("request")
			return nil
		},
	}))

	// Setup custom context on all routes
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.Context{Context: c}
			cc.SetDatabaseClient(db)
			return next(cc)
		}
	})

	e.Any("/", staticPage(pages.Index()))
	e.Any("/error/", withContext(handlers.Error))
	e.Any("/login/", withContext(handlers.Login))
	e.Any("/logout/", withContext(handlers.Logout))
	e.POST("/login/google/redirect/", withContext(handlers.GoogleAuthRedirect))

	e.Any("/cookie-policy/", staticPage(pages.CookiePolicy()))
	e.Any("/privacy-policy/", staticPage(pages.PrivacyPolicy()))
	e.Any("/terms-of-service/", staticPage(pages.TermsOfService()))

	appGroup := e.Group("/app", sessionCheckMiddleware(db))
	appGroup.Any("/", withContext(app.Home))
	appGroup.GET("/group/:groupId/", withContext(app.Group))
	appGroup.GET("/group/:groupId/manage/", withContext(app.ManageGroup))
	appGroup.GET("/group/:groupId/list/:listId/", withContext(app.List))
	appGroup.GET("/group/:groupId/list/:listId/manage/", withContext(app.ManageList))
	appGroup.GET("/settings/", withContext(app.Settings))

	apiGroup := e.Group("/api", sessionCheckMiddleware(db))
	apiGroup.POST("/groups/", withContext(api.CreateGroup))
	apiGroup.POST("/groups/:groupId/lists/", withContext(api.CreateList))
	apiGroup.POST("/groups/:groupId/lists/:listId/items/", withContext(api.CreateItem))
	apiGroup.PUT("/groups/:groupId/lists/:listId/items/:itemId/checked/", withContext(api.ItemSetChecked))
	apiGroup.POST("/groups/:groupId/invites/", withContext(api.CreateGroupInvite))
	apiGroup.POST("/groups/invites/redeem/", withContext(api.RedeemGroupInvite))

	e.Any("/*", pageNotFound)
	return e
}

func pageNotFound(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Page "+c.Request().URL.Path+" does not exist")
}

// Convert a custom context handler into an echo handler
func withContext(h func(*handlers.Context) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		return h(c.(*handlers.Context))
	}
}

// Create an echo handler from Page
func staticPage(page templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*handlers.Context)
		return cc.RenderPage(page)
	}
}
