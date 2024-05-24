package api

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/prisma/db"
)

type GroupCreateForm struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

func CreateGroup(c *handlers.Context) error {
	var data GroupCreateForm
	if err := c.Bind(&data); err != nil {
		// Respond for HTMX
	}

	// Create a group
	group, err := c.DB().Group.CreateOne(db.Group.Name.Set(data.Name), db.Group.Desc.Set(data.Description)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a group&context="+err.Error())
	}

	// Create a membership
	_, err = c.DB().GroupMember.CreateOne(db.GroupMember.Group.Link(db.Group.ID.Equals(group.ID)), db.GroupMember.User.Link(db.User.ID.Equals(c.User().ID)), db.GroupMember.Permissions.Set("v0/0")).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a group&context="+err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/app/group/"+group.ID)
}
