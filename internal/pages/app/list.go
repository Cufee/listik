package app

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	"github.com/cufee/shopping-list/internal/pages"
	g "github.com/maragudk/gomponents"
)

func List(id string) (pages.Page, error) {
	return pages.NewPage(
		bulma.Content(bulma.None(),
			g.Text("list: "+id),
		),
		pages.WithTitle("List"),
	)
}
