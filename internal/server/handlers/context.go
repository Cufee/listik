package handlers

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/pages"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func (c *Context) Render(p pages.Page, err error) error {
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=We ran into an error&context="+err.Error())
	}
	return p.Node(c.Path()).Render(c.Response().Writer)
}
