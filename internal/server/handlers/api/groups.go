package api

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

type GroupCreateForm struct {
	Name        string `form:"name"`
	Description string `form:"description"`
}

func CreateGroup(c *handlers.Context) error {
	var data GroupCreateForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a group&context="+err.Error())
	}

	if len(data.Name) < 5 || len(data.Name) > 21 {
		return c.RenderPartial(app.CreateGroupDialog(true, map[string]string{"name": data.Name, "description": data.Description}, map[string]string{"name": "group name should be between 5 and 21 characters"}))
	}
	if len(data.Description) > 80 {
		return c.RenderPartial(app.CreateGroupDialog(true, map[string]string{"name": data.Name, "description": data.Description}, map[string]string{"description": "group description is limited to 80 characters"}))
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

func RedeemGroupInvite(c *handlers.Context) error {
	return c.RenderPartial(app.OnboardingGroups(map[string]string{"invite-code": c.FormValue("invite-code")}, map[string]string{"invite-code": "invalid invite code"}))
}
