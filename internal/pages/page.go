package pages

import (
	"strings"

	"github.com/cufee/shopping-list/internal/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func Page(title, path string, body g.Node) g.Node {
	// HTML5 boilerplate document
	return html5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			h.Link(h.Href("https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css"), h.Rel("stylesheet")),
			h.Script(h.Src("https://unpkg.com/htmx.org@1.9.12")),
		},
		Body: []g.Node{
			g.If(strings.HasPrefix(path, "/app"), components.AppNavbar(path)),
			g.If(!strings.HasPrefix(path, "/app"), components.Navbar(path)),
			components.Container(
				components.Section(body),
				PageFooter(),
			),
		},
	})
}

func html5(p c.HTML5Props) g.Node {
	return h.Doctype(
		h.HTML(g.If(p.Language != "", h.Lang(p.Language)),
			h.DataAttr("theme", "dark"),
			h.Head(
				h.Meta(h.Charset("utf-8")),
				h.Meta(h.Name("viewport"), h.Content("width=device-width, initial-scale=1")),
				h.TitleEl(g.Text(p.Title)),
				g.If(p.Description != "", h.Meta(h.Name("description"), h.Content(p.Description))),
				g.Group(p.Head),
			),
			h.Body(g.Group(p.Body)),
		),
	)
}

func PageFooter() g.Node {
	return h.Footer()
}
