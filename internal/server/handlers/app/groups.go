package app

import (
	"net/http"
	"time"

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

	group, err := c.DB().Group.FindUnique(db.Group.ID.Equals(member.GroupID)).With(db.Group.Lists.Fetch()).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.Page(http.StatusOK, app.Group{Group: group, Lists: group.Lists()}.Render())
}

func ManageGroup(c *handlers.Context) error {
	member, err := c.DB().GroupMember.FindFirst(db.GroupMember.UserID.Equals(c.User().ID), db.GroupMember.GroupID.Equals(c.Param("groupId"))).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	group, err := c.DB().Group.FindUnique(db.Group.ID.Equals(member.GroupID)).With(
		db.Group.Invites.Fetch(db.GroupInvite.ExpiresAt.After(time.Now())),
		db.Group.Members.Fetch(),
		db.Group.Lists.Fetch(),
	).Exec(c.Request().Context())
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	var memberUserIDs []string
	for _, member := range group.Members() {
		memberUserIDs = append(memberUserIDs, member.UserID)
	}

	memberUsers, err := c.DB().User.FindMany(db.User.ID.In(memberUserIDs)).Exec(c.Request().Context())
	if err != nil && !db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.Page(http.StatusOK, app.ManageGroup{Group: group, Lists: group.Lists(), Members: memberUsers, Invites: group.Invites()}.Render())
}
