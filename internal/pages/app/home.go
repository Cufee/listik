package app

import (
	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/components/bulma"
	"github.com/cufee/shopping-list/internal/pages"
	"github.com/cufee/shopping-list/prisma/db"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

type HomePageProps struct {
	Lists       map[string]db.ListModel
	RecentLists []string
}

func Home(props HomePageProps) (pages.Page, error) {
	var listsSlice []db.ListModel
	for _, list := range props.Lists {
		listsSlice = append(listsSlice, list)
	}

	return pages.NewPage(
		bulma.Content(bulma.None(),
			g.If(len(props.RecentLists) > 0,
				bulma.Content(bulma.None(),
					h.Div(
						h.H3(g.Text("Recent")),
					),
					bulma.Grid(bulma.None(),
						g.Map(props.RecentLists, func(id string) g.Node {
							return bulma.Cell(bulma.None(), components.ListCard(props.Lists[id]))
						})...,
					),
				),
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
				bulma.Grid(bulma.None(),
					g.Map(listsSlice, func(list db.ListModel) g.Node {
						return bulma.Cell(bulma.None(), components.ListCard(list))
					})...,
				),
			),
		),
	)
}
