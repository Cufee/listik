package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/cufee/shopping-list/internal/templates/pages"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context

	authenticated bool
	user          *db.UserModel
	db            *db.PrismaClient
}

/*
Renders a page into response writer
  - Adds a navbar and footer based on the current page context
*/
func (c *Context) RenderPage(page templ.Component) error {
	return pages.Wrapper(c.Path(), c.Authenticated(), page).Render(c.Request().Context(), c.Response().Writer)
}

func (c *Context) Redirect(code int, path string) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		c.Response().Header().Set("HX-Redirect", path)
		return c.String(http.StatusOK, "")
	}
	return c.Context.Redirect(code, path)
}

func (c *Context) SetUser(user *db.UserModel) {
	if user != nil {
		c.authenticated = true
		c.user = user
	}
}

func (c *Context) User() *db.UserModel {
	if c.authenticated {
		return c.user
	}
	return nil
}

func (c *Context) Authenticated() bool {
	return c.authenticated && c.user != nil
}

func (c *Context) SetDatabaseClient(client *db.PrismaClient) {
	c.db = client
}

func (c *Context) DB() *db.PrismaClient {
	return c.db
}
