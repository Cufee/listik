package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type template struct{}

func (t *template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	component, ok := data.(templ.Component)
	if !ok {
		return c.Redirect(http.StatusInternalServerError, fmt.Sprintf("/error?message=Failed to display %s page&context=invalid component", c.Path()))
	}
	return component.Render(c.Request().Context(), w)
}
