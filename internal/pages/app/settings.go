package app

import (
	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/pages"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Settings() (pages.Page, error) {
	return pages.NewPage(
		components.Container(
			h.A(
				h.Href("/"),
				g.Text("Log Out"),
				h.Class("button is-danger"),
			),
		),
		pages.WithTitle("Settings"),
	)
}
