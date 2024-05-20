package pages

import (
	"github.com/cufee/shopping-list/internal/components"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func About() (Page, error) {
	return NewPage(
		components.Container(
			h.H1(g.Text("About this site")),
			h.P(g.Text("This is a site showing off gomponents.")),
		),
		WithTitle("About"),
	)
}
