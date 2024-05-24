package handlers

import "github.com/cufee/shopping-list/internal/templates/pages"

func Error(c *Context) error {
	return c.RenderPage(pages.Error(c.QueryParam("message"), c.QueryParam("context")))
}
