package app

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Settings() (string, g.Node) {
	return "Settings", h.Div(
		h.A(
			h.Href("/"),
			g.Text("Log Out"),
			h.Class("button is-danger"),
		),
	)
}
