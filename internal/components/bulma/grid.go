package bulma

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Grid(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "grid")
}

func Cell(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "cell")
}
