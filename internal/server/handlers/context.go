package handlers

import (
	"net/http"
	"strings"

	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/pages"
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
func (c *Context) RenderPage(p pages.Page, err error) error {
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=We ran into an error&context="+err.Error())
	}

	// TODO: This needs to be done better, building a navbar based on context should be more flexible and cleaner
	if strings.HasPrefix(c.Path(), "/app") {
		p.SetOption(pages.WithNavbar(components.AppNavbar(c.Path())))
	} else {
		p.SetOption(pages.WithNavbar(components.Navbar(c.Path())))
	}
	p.SetOption(pages.WithFooter(components.Footer()))

	return p.Node(c.Path()).Render(c.Response().Writer)
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
