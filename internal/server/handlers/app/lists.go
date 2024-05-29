package app

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

func List(c *handlers.Context) error {
	member, err := c.DB().GroupMember.FindFirst(db.GroupMember.UserID.Equals(c.User().ID), db.GroupMember.GroupID.Equals(c.Param("groupId"))).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=list not found&context="+err.Error())
	}

	list, err := c.DB().List.FindFirst(db.List.ID.Equals(c.Param("listId")), db.List.GroupID.Equals(member.GroupID)).With(db.List.Items.Fetch()).With(db.List.Group.Fetch()).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app/group/"+member.GroupID)
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.Page(http.StatusOK, app.List{Group: list.Group(), List: list, Items: list.Items()}.Render())
}

func ManageList(c *handlers.Context) error {
	member, err := c.DB().GroupMember.FindFirst(db.GroupMember.UserID.Equals(c.User().ID), db.GroupMember.GroupID.Equals(c.Param("groupId"))).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=list not found&context="+err.Error())
	}

	list, err := c.DB().List.FindFirst(db.List.ID.Equals(c.Param("listId")), db.List.GroupID.Equals(member.GroupID)).With(db.List.Items.Fetch()).With(db.List.Group.Fetch()).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app/group/"+member.GroupID)
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.Page(http.StatusOK, app.ManageList{Group: list.Group(), List: list, Items: list.Items()}.Render())
}
