package pages

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func About() (string, g.Node) {
	return "About", h.Div(
		h.H1(g.Text("About this site")),
		h.P(g.Text("This is a site showing off gomponents.")),
	)
}
