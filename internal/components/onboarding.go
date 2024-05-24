package components

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	g "github.com/maragudk/gomponents"
)

func OnboardingActionCard(title string, body g.Node) g.Node {
	return bulma.Card(bulma.None(),
		bulma.CardHeader(bulma.None(),
			bulma.CardHeaderTitle(bulma.Class("is-centered"),
				g.Text(title),
			),
		),
		bulma.CardContent(bulma.Class("is-flex", "is-flex-direction-column", "is-justify-content-center"),
			body,
		),
	)
}
