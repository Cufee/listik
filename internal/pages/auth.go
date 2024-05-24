package pages

import (
	"os"

	"github.com/cufee/shopping-list/internal/components/bulma"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Login() (Page, error) {
	return NewPage(
		bulma.Content(bulma.None(),
			h.H1(g.Text("Login page")),
			h.A(
				h.Href("/app"),
				g.Text("Log In"),
				h.Class("button is-primary"),
			),
		),
		WithTitle("Login"),
	)
}

func SignUp() (Page, error) {
	return NewPage(
		bulma.Content(bulma.None(),
			h.H1(g.Text("Create an account")),
			h.A(
				h.Href("/app"),
				g.Text("Log In"),
				h.Class("button is-primary"),
			),
		),
		WithTitle("New user? - Sign Up!"),
	)
}

func Logout() (Page, error) {
	return NewPage(
		bulma.Content(bulma.None(),
			h.H1(g.Text("You were logged out of "+os.Getenv("APP_NAME"))),
		),
		WithTitle("Logged Out"),
	)
}
