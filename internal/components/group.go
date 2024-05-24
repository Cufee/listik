package components

import (
	"github.com/cufee/shopping-list/internal/components/bulma"
	"github.com/cufee/shopping-list/prisma/db"
	g "github.com/maragudk/gomponents"
)

func GroupCard(group *db.GroupModel) g.Node {
	return bulma.Content(bulma.None(),
		g.Text("group: "+group.ID),
	)
}
