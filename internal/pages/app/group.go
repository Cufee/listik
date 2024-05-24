package app

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	"github.com/cufee/shopping-list/internal/pages"
	"github.com/cufee/shopping-list/prisma/db"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Group(group *db.GroupModel) (pages.Page, error) {
	return pages.NewPage(
		bulma.Content(bulma.None(),
			bulma.Content(bulma.None(),
				h.Div(
					h.H3(g.Text("Recent")),
				),
				bulma.Grid(bulma.None()),
			),

			bulma.Buttons(bulma.Class("is-centered"),
				bulma.Button(bulma.Class("is-primary"),
					g.Text("New List"),
				),
			),

			bulma.Content(bulma.None(),
				h.Div(
					h.H3(g.Text("All")),
				),
				bulma.Grid(bulma.None()),
			),
		),
	)
}
