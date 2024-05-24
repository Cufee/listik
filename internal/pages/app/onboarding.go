package app

import (
	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/components/bulma"
	"github.com/cufee/shopping-list/internal/pages"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func OnboardingGroups() (pages.Page, error) {
	return pages.NewPage(
		bulma.Content(bulma.None(),
			h.Div(
				bulma.Title(bulma.Class("has-text-centered"),
					g.Text("You are not a part of a group yet"),
				),
				bulma.Subtitle(bulma.Class("has-text-centered"),
					g.Text("Start by making your own or using a invite code to join an existing group!"),
				),
			),

			bulma.Grid(bulma.None(),
				bulma.Cell(bulma.None(),
					components.OnboardingActionCard("Join a group",
						bulma.Button(bulma.Class("is-primary"), g.Text("Join")),
					)),
				bulma.Cell(bulma.None(),
					components.OnboardingActionCard("Create a new group",
						bulma.Button(bulma.Class("is-primary"), g.Text("Create")),
					)),
			),
		),

		pages.WithTitle("Welcome!"),
	)
}

func OnboardingLists() (pages.Page, error) {
	return pages.NewPage(
		bulma.Content(bulma.None(),
			g.Text("create a first list flow"),
		),
		pages.WithTitle("Welcome!"),
	)
}
