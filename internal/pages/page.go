package pages

import (
	"os"
	"strings"

	"github.com/cufee/shopping-list/internal/components"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

type Page struct {
	options options
	body    g.Node
}

type options struct {
	title string
	head  []g.Node
}

var defaultOptions = options{
	title: os.Getenv("APP_NAME"),
}

type Option func(op *options) error

func WithTitle(title string) Option {
	return func(op *options) error {
		op.title += " - " + title
		return nil
	}
}

func NewPage(body g.Node, opts ...Option) (Page, error) {
	options := defaultOptions
	for _, o := range opts {
		err := o(&options)
		if err != nil {
			return Page{}, err
		}
	}
	return Page{
		options: options,
		body:    body,
	}, nil
}

func (p Page) Node(path string) g.Node {
	// HTML5 boilerplate document
	return html5(c.HTML5Props{
		Title:    p.options.title,
		Language: "en",
		Head: append([]g.Node{
			h.Link(h.Href("https://cdn.jsdelivr.net/npm/bulma@1.0.0/css/bulma.min.css"), h.Rel("stylesheet")),
			h.Script(h.Src("https://unpkg.com/htmx.org@1.9.12")),
		}, p.options.head...),
		Body: []g.Node{
			g.If(strings.HasPrefix(path, "/app"), components.AppNavbar(path)),
			g.If(!strings.HasPrefix(path, "/app"), components.Navbar(path)),
			components.Container(
				components.Section(p.body),
				Footer(),
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

func Footer() g.Node {
	return h.Footer()
}
