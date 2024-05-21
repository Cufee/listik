package components

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func ListCard() g.Node {
	return h.Div(
		h.Class("card has-background-black-ter"),
		h.Header(
			h.Class("card-header"),
			h.P(
				h.Class("card-header-title"),
				g.Text("List Name"),
			),
		),
		h.Div(
			h.Class("card-content p-4"),
			h.Div(
				// h.Class("content"),
				h.P(g.Text("Some Text")),
				h.Time(
					g.Attr("datetime", "2016-1-1"),
					g.Text("11:09 PM - 1 Jan 2016"),
				),
			),
		),
		h.Footer(
			h.Class("card-footer"),
			h.A(
				h.Href("#"),
				h.Class("card-footer-item"),
				g.Text("Action"),
			),
		),
	)
}
