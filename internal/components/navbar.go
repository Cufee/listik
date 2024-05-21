package components

import (
	"os"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
)

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string) g.Node {
	links := []PageLink{
		{Path: "/about/", Name: "About"},
	}

	return b.Navbar(
		hx.Boost("true"),
		branding("/", currentPath == "/"),
		b.NavbarStart(
			h.Class("is-active"),
			g.Group(g.Map(links, func(l PageLink) g.Node {
				return NavbarLink(l.Path, l.Name, currentPath == l.Path)
			})),
		),
		b.NavbarEnd(
			b.NavbarItem(
				b.Buttons(
					h.A(
						h.Href("/sign-up"),
						h.Class("button is-primary is-small"),
						h.Strong(g.Text("Sign Up")),
					),
					h.A(
						h.Href("/login"),
						h.Class("button is-light is-small"),
						h.Strong(g.Text("Log In")),
					),
				),
			),
		),
	)
}

func AppNavbar(currentPath string) g.Node {
	links := []PageLink{
		{Path: "/app/settings/", Name: "Settings"},
	}

	return b.Navbar(
		hx.Boost("true"),
		branding("/app", currentPath == "/app/"),
		b.NavbarEnd(
			h.Class("is-active"),
			g.Group(g.Map(links, func(l PageLink) g.Node {
				return NavbarLink(l.Path, l.Name, currentPath == l.Path)
			})),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return h.A(
		h.Href(path),
		c.Classes{
			"is-active":   active,
			"navbar-item": true,
		},
		g.Text(text),
	)
}

func branding(href string, highlight bool) any {
	return b.NavbarBrand(
		h.A(
			h.Href(href),
			c.Classes{
				"is-active":   highlight,
				"navbar-item": true,
			},

			Icon(20),
			h.Span(
				h.Class("pl-1"),
				h.Strong(g.Text(os.Getenv("APP_NAME"))),
			),
		),
	)
}
