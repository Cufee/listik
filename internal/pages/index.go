package pages

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Index() (Page, error) {
	return NewPage(
		bulma.Content(bulma.None(),
			h.H1(g.Text("Welcome to this example page")),
			h.P(g.Text("I hope it will make you happy. 😄 It's using Bulma for styling.")),
		),
	)
}
