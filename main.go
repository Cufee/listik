package main

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/pages"
	"github.com/cufee/shopping-list/internal/pages/app"
	g "github.com/maragudk/gomponents"
)

func main() {
	http.Handle("/", createHandler(pages.Index()))
	http.Handle("/about", createHandler(pages.About()))

	http.Handle("/app", createHandler(app.Home()))
	http.Handle("/app/settings", createHandler(app.Settings()))

	http.Handle("/login", createHandler(pages.Index()))
	http.Handle("/sign-up", createHandler(pages.Index()))

	panic(http.ListenAndServe("localhost:8081", nil))
}

func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Rendering a Node is as simple as calling Render and passing an io.Writer
		_ = pages.Page(title, r.URL.Path, body).Render(w)
	}
}
