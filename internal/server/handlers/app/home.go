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
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to retrieve your groups&context="+err.Error())
	}

	// If this user has no groups, they should go through the onboarding flow
	if len(memberships) == 0 || c.QueryParam("onboarding") == "true" {
		return c.RenderPage(app.OnboardingGroups(nil, nil))
	}

	var groupIDs []string
	for _, m := range memberships {
		groupIDs = append(groupIDs, m.GroupID)
	}

	groups, err := c.DB().Group.FindMany(db.Group.ID.In(groupIDs)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to retrieve your groups&context="+err.Error())
	}

	return c.RenderPage(app.Home{Groups: groups}.Render())
}

func Settings(c *handlers.Context) error {
	// id := c.Param("id")
	return c.RenderPage(app.Settings())
}
