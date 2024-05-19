package components

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Container(children ...g.Node) g.Node {
	return h.Div(h.Class("content"), g.Group(children))
}

func Section(children ...g.Node) g.Node {
	return h.Div(h.Class("section"), g.Group(children))
}
