package pages

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Error(message string, context string) (Page, error) {
	if message == "" {
		message = "We are not sure what happened"
	}

	return NewPage(
		bulma.Content(bulma.None(),
			h.Div(
				h.Class("notification is-danger"),
				h.Strong(g.Text("Something went wrong")),
				h.P(g.Text(message)),
				g.If(context != "",
					h.P(g.Text(context)),
				),
			),
		),
		WithTitle("Error"),
	)
}
