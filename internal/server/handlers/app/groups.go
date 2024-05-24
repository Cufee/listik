package app

import (
	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
)

func GroupOverview(c *handlers.Context) error {
	_ = c.Param("id")
	return c.RenderPage(app.Home())
}

func GroupLists(c *handlers.Context) error {
	_ = c.Param("id")
	return c.RenderPage(app.Home())
}

func List(c *handlers.Context) error {
	_ = c.Param("id")
	return c.RenderPage(app.Home())
}
