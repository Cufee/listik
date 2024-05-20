package handlers

import (
	"github.com/cufee/shopping-list/internal/pages/app"
)

func ViewList(c *Context) error {
	id := c.Param("id")
	return c.Render(app.List(id))
}
