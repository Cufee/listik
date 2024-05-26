package server

import (
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
	"github.com/cufee/shopping-list/prisma/db"

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

	// Setup custom context on all routes
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &handlers.Context{Context: c}
			cc.SetDatabaseClient(db)
			return next(cc)
		}
	})

	e.GET("/", staticPage(pages.Index()))
	e.GET("/error/", withContext(handlers.Error))
	e.GET("/login/", staticPage(pages.Login()))

	appGroup := e.Group("/app", sessionCheckMiddleware(db))
	appGroup.GET("/", withContext(app.Home))
	appGroup.GET("/group/:groupId/", withContext(app.Group))
	appGroup.GET("/group/:groupId/list/:listId/", withContext(app.List))
	appGroup.GET("/settings/", withContext(app.Settings))

	apiGroup := e.Group("/api", sessionCheckMiddleware(db))
	apiGroup.POST("/groups/", withContext(api.CreateGroup))
	apiGroup.POST("/groups/:groupId/lists/", withContext(api.CreateList))
	apiGroup.POST("/groups/:groupId/lists/:listId/items/", withContext(api.CreateItem))
	apiGroup.PUT("/groups/:groupId/lists/:listId/items/:itemId/checked/", withContext(api.ItemSetChecked))

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
func staticPage(page templ.Component) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*handlers.Context)
		return cc.RenderPage(page)
	}
}
