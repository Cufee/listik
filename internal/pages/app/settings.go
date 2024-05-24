package app

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	"github.com/cufee/shopping-list/internal/pages"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	h "github.com/maragudk/gomponents/html"
)

func Settings() (pages.Page, error) {
	return pages.NewPage(
		bulma.Content(bulma.None(),
			h.A(
				hx.Boost("true"),
				h.Href("/"),
				g.Text("Log Out"),
				h.Class("button is-danger"),
			),
		),
		pages.WithTitle("Settings"),
	)
}
