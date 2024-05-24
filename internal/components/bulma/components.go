package bulma

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Card(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card")
}

func CardHeader(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card-header")
}

func CardHeaderTitle(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.P, "card-header-title")
}

func CardHeaderIcon(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card-header-icon")
}

func CardImage(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card-image")
}

func CardContent(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card-content")
}

func CardFooter(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card-footer")
}

func CardFooterItem(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "card-footer-item")
}

func Navbar(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "navbar")
}

func NavbarBrand(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "navbar-brand")
}

func NavbarStart(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "navbar-start")
}

func NavbarEnd(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "navbar-end")
}

func NavbarItem(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Div, "navbar-item")
}

func NavbarLink(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.A, "navbar-link")
}

func NavbarDivider(opts Options, children ...g.Node) g.Node {
	return nodeWithBaseClass(opts, children, h.Span, "navbar-divider")
}
