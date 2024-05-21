package server

import (
	"io/fs"
	"net/http"

	"github.com/cufee/shopping-list/internal/pages"
	"github.com/cufee/shopping-list/internal/pages/app"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Create a new echo.Echo instance with all routes registered
func New(assets fs.FS) *echo.Echo {
	e := echo.New()
	e.StaticFS("/static", assets)

	e.Pre(middleware.AddTrailingSlash()) // echo does not route correctly when a `/` route is attached to a group, this is to fix that issue

	e.GET("/", fromPage(pages.Index))
	e.GET("/error/", withContext(handlers.Error))
	e.GET("/about/", fromPage(pages.About))
	e.GET("/login/", fromPage(pages.Login))
	e.GET("/sign-up/", fromPage(pages.SignUp))

	eApp := e.Group("/app")
	eApp.GET("/", fromPage(app.Home))
	eApp.GET("/settings/", fromPage(app.Settings))

	eApp.GET("/list/", redirect("/app"))
	eApp.GET("/list/:id/", withContext(handlers.ViewList))

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
		cc := &handlers.Context{Context: c}
		return h(cc)
	}
}

// Create an echo handler from Page
func fromPage(handler func() (pages.Page, error)) echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := handler()
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to load this page&context="+err.Error())
		}
		return page.Node(c.Path()).Render(c.Response().Writer)
	}
}
