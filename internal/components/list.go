package components

import (
	"fmt"

	"github.com/cufee/shopping-list/prisma/db"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func ListCard(list db.ListModel) g.Node {
	description, hasDescription := list.Desc()

	return h.Div(
		h.Class("card has-background-black-ter"),
		h.Header(
			h.Class("card-header"),
			h.P(
				h.Class("card-header-title"),
				g.Text(list.Name),
			),
		),
		h.Div(
			h.Class("card-content p-4"),
			h.Div(
				// h.Class("content"),
				g.If(hasDescription, h.P(g.Text(description))),
				h.Time(
					g.Attr("datetime", "2016-1-1"),
					g.Text("11:09 PM - 1 Jan 2016"),
				),
			),
		),
		h.Footer(
			h.Class("card-footer"),
			h.A(
				h.Href(fmt.Sprintf("/app/list/%s", list.ID)),
				h.Class("card-footer-item"),
				g.Text("View"),
			),
		),
	)
}
