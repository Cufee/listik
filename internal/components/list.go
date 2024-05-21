package components

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
	b "github.com/willoma/bulma-gomponents"
)

func ListCard() g.Node {
	return b.Card(
		b.CardHeader(
			b.CardHeaderTitle(
				g.Text("List Name"),
			),
		),
		h.Div(
			g.Text("Some Text"),
		),
		b.CardFooter(
			h.A(
				h.Href("#"),
				h.Class("card-footer-item"),
				g.Text("Action"),
			),
		),
	)
}
