package server

import (
	"io/fs"
	"net/http"

	"github.com/cufee/shopping-list/internal/pages"
	"github.com/cufee/shopping-list/prisma/db"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/server/handlers/app"
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

	// Setup custom context on all routes
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.Context{Context: c}
			cc.SetDatabaseClient(db)
			return next(cc)
		}
	})

	e.GET("/", staticPage(pages.Index))
	e.GET("/error/", withContext(handlers.Error))
	e.GET("/about/", staticPage(pages.About))
	e.GET("/login/", staticPage(pages.Login))
	e.GET("/sign-up/", staticPage(pages.SignUp))

	eApp := e.Group("/app")
	eApp.Use(sessionCheckMiddleware(db))

	eApp.GET("/", withContext(app.Home))

	eApp.GET("/:groupId/", withContext(app.GroupOverview))
	eApp.GET("/:groupId/list/:listId/", withContext(app.List))

	eApp.GET("/settings/", withContext(app.Settings))

	e.GET("/*", pageNotFound)
	return e
}

func pageNotFound(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Page "+c.Request().URL.Path+" does not exist")
}

// Create a temp redirect handler for path
func redirect(path string) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, path)
	}
}

// Convert a custom context handler into an echo handler
func withContext(h func(*handlers.Context) error) func(c echo.Context) error {
	return func(c echo.Context) error {
		return h(c.(*handlers.Context))
	}
}

// Create an echo handler from Page
func staticPage(handler func() (pages.Page, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*handlers.Context)
		return cc.RenderPage(handler())
	}
}
