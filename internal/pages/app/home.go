package app

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Home() (string, g.Node) {
	return "App", h.Div(
		h.H1(g.Text("App Home Page")),
		h.P(g.Text("This is a site showing off gomponents.")),
	)
}
