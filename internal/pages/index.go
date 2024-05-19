package pages

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Index() (string, g.Node) {
	return "Welcome!", h.Div(
		h.H1(g.Text("Welcome to this example page")),
		h.P(g.Text("I hope it will make you happy. 😄 It's using Bulma for styling.")),
	)
}
