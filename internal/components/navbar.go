package components

import (
	"os"

	"github.com/cufee/shopping-list/internal/components/bulma"
	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string) g.Node {
	links := []PageLink{
		{Path: "/about/", Name: "About"},
	}

	return bulma.Navbar(bulma.With(hx.Boost("true")),
		branding("/", currentPath == "/"),
		bulma.NavbarStart(bulma.Class("is-active"),
			g.Group(g.Map(links, func(l PageLink) g.Node {
				return navbarLink(l.Path, l.Name, currentPath == l.Path)
			})),
		),
		bulma.NavbarEnd(bulma.None(),
			bulma.NavbarItem(bulma.None(),
				bulma.Buttons(bulma.None(),
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

	return bulma.Navbar(bulma.With(hx.Boost("true")),
		branding("/app", currentPath == "/app/"),
		bulma.NavbarEnd(bulma.Class("is-active"),
			g.Group(g.Map(links, func(l PageLink) g.Node {
				return navbarLink(l.Path, l.Name, currentPath == l.Path)
			})),
		),
	)
}

// NavbarLink is a link in the Navbar.
func navbarLink(path, text string, active bool) g.Node {
	return h.A(
		h.Href(path),
		c.Classes{
			"is-active":   active,
			"navbar-item": true,
		},
		g.Text(text),
	)
}

func branding(href string, highlight bool) g.Node {
	return bulma.NavbarBrand(bulma.None(),
		bulma.A(bulma.With(h.Href(href)).Class("navbar-item").ClassIf("is-active", highlight),
			Icon(20),
			h.Span(
				h.Class("pl-1"),
				h.Strong(g.Text(os.Getenv("APP_NAME"))),
			),
		),
	)

}
