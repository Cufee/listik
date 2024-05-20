package app

import (
	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/pages"
	g "github.com/maragudk/gomponents"
)

func List(id string) (pages.Page, error) {
	return pages.NewPage(
		components.Container(
			g.Text("list: "+id),
		),
		pages.WithTitle("List"),
	)
}
