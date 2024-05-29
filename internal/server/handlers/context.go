package handlers

import (
	"errors"
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
func (c *Context) Page(code int, node templ.Component) error {
	return c.Context.Render(code, "", pages.Wrapper(c.Path(), c.Authenticated(), node))
}

/*
Renders a single component into response writer
*/
func (c *Context) Partial(code int, node templ.Component) error {
	return c.Context.Render(code, "", node)
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

func (c *Context) Member(groupID string) (*db.GroupMemberModel, error) {
	if !c.Authenticated() {
		return nil, errors.New("not authenticated")
	}
	return c.DB().GroupMember.FindFirst(db.GroupMember.UserID.Equals(c.User().ID), db.GroupMember.GroupID.Equals(groupID)).Exec(c.Request().Context())
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
