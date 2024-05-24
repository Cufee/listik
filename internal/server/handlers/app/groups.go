package app

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

func Group(c *handlers.Context) error {
	member, err := c.DB().GroupMember.FindFirst(db.GroupMember.UserID.Equals(c.User().ID), db.GroupMember.GroupID.Equals(c.Param("groupId"))).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	group, err := c.DB().Group.FindUnique(db.Group.ID.Equals(member.GroupID)).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	lists, err := c.DB().List.FindMany(db.List.GroupID.Equals(group.ID)).With(db.List.Items.Fetch()).Exec(c.Request().Context())
	if err != nil && !db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.RenderPage(app.Group{Group: group, Lists: lists}.Render())
}

func List(c *handlers.Context) error {
	member, err := c.DB().GroupMember.FindFirst(db.GroupMember.UserID.Equals(c.User().ID), db.GroupMember.GroupID.Equals(c.Param("groupId"))).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=list not found&context="+err.Error())
	}

	list, err := c.DB().List.FindFirst(db.List.ID.Equals(c.Param("listId")), db.List.GroupID.Equals(member.GroupID)).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app/group/"+member.GroupID)
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.RenderPage(app.List{List: list}.Render())
}
