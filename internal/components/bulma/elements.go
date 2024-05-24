package bulma

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func A(opts Options, children ...g.Node) g.Node {
	return h.A(opts.fromNodes(children)...)
}

func Button(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "button")
}

func Buttons(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "buttons")
}

func Content(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "content")
}

func Section(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "section")
}

func Title(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.P, "title")
}

func Subtitle(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.P, "subtitle")
}
