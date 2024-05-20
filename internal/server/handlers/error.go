package handlers

import "github.com/cufee/shopping-list/internal/pages"

func Error(c *Context) error {
	return c.Render(pages.Error(c.QueryParam("message"), c.QueryParam("context")))
}
