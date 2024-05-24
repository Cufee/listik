package app

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

func Home(c *handlers.Context) error {
	memberships, err := c.DB().GroupMember.FindMany(db.GroupMember.UserID.Equals(c.User().ID)).Exec(c.Request().Context())
	if err != nil && !db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=Failed to retrieve your groups&context="+err.Error())
	}

	// If this user has no groups, they should go through the onboarding flow
	if len(memberships) == 0 {
		return c.RenderPage(app.Home())
	}

	// id := c.Param("id")
	return c.RenderPage(app.Home())
}

func Settings(c *handlers.Context) error {
	// id := c.Param("id")
	// return c.RenderPage(app.Settings())
	return nil
}
