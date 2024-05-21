package app

import (
	"github.com/cufee/shopping-list/internal/components"
	"github.com/cufee/shopping-list/internal/pages"
	b "github.com/willoma/bulma-gomponents"
)

func Home() (pages.Page, error) {
	return pages.NewPage(
		b.Grid(
			b.Cell(components.ListCard()),
			b.Cell(components.ListCard()),
			b.Cell(components.ListCard()),
			b.Cell(components.ListCard()),
		),
	)
}
