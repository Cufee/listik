package app

import (
	"github.com/cufee/shopping-list/internal/pages/app"
	"github.com/cufee/shopping-list/internal/server/handlers"
)

func GroupOverview(c *handlers.Context) error {
	id := c.Param("id")
	return c.RenderPage(app.List(id))
}

func GroupLists(c *handlers.Context) error {
	id := c.Param("id")
	return c.RenderPage(app.List(id))
}

func List(c *handlers.Context) error {
	id := c.Param("id")
	return c.RenderPage(app.List(id))
}
