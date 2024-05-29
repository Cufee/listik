package handlers

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/templates/pages"
)

func Error(c *Context) error {
	return c.Page(http.StatusOK, pages.Error(c.QueryParam("message"), c.QueryParam("context")))
}
