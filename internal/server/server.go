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
	e.Renderer = &template{}

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

	eApp := e.Group("/app", sessionCheckMiddleware(db))

	eApp.Any("/", withContext(app.Home))
	eApp.Any("/group/:groupId/", withContext(app.Group))
	eApp.Any("/group/:groupId/manage/", withContext(app.ManageGroup))
	eApp.Any("/group/:groupId/list/:listId/", withContext(app.List))
	eApp.Any("/settings/", withContext(app.Settings))

	eApi := e.Group("/api", sessionCheckMiddleware(db))

	apiGroups := eApi.Group("/groups")
	apiGroups.POST("/", withContext(api.CreateGroup))
	apiGroups.POST("/invites/redeem/", withContext(api.RedeemGroupInvite))
	apiGroups.POST("/:groupId/invites/", withContext(api.CreateGroupInvite))

	apiTags := apiGroups.Group("/:groupId/tags")
	apiTags.POST("/", withContext(api.CreateItemTag))
	apiTags.PATCH("/:tagId/", withContext(api.CreateItemTag))
	apiTags.DELETE("/:tagId/", withContext(api.CreateItemTag))

	apiLists := apiGroups.Group("/:groupId/lists")
	apiLists.POST("/", withContext(api.CreateList))
	apiLists.POST("/:listId/items/", withContext(api.CreateItem))
	apiLists.PATCH("/:listId/complete/", withContext(api.ListSetComplete))
	apiLists.DELETE("/:listId/items/:itemId/", withContext(api.DeleteItem))
	apiLists.PUT("/:listId/items/:itemId/checked/", withContext(api.ItemSetChecked))

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
		return cc.Page(http.StatusOK, page)
	}
}
