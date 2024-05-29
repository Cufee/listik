package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/componenets/group"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
	"github.com/rs/zerolog/log"
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

	if len(data.Name) < 1 || len(data.Name) > 80 {
		return c.Partial(http.StatusUnprocessableEntity, app.CreateGroupDialog(true, makeInputsMap(data), map[string]string{"name": "group name should be between 1 and 80 characters"}))
	}
	if len(data.Description) > 80 {
		return c.Partial(http.StatusUnprocessableEntity, app.CreateGroupDialog(true, makeInputsMap(data), map[string]string{"description": "group description is limited to 80 characters"}))
	}

	// Create a group
	group, err := c.DB().Group.CreateOne(db.Group.Owner.Link(db.User.ID.Equals(c.User().ID)), db.Group.Name.Set(data.Name), db.Group.Desc.Set(data.Description)).Exec(c.Request().Context())
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

type GroupInviteCreateForm struct {
	GroupID string `param:"groupId"`
}

func CreateGroupInvite(c *handlers.Context) error {
	var data GroupInviteCreateForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create an invite&context="+err.Error())
	}

	// Check if a user belong to this group
	member, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create an invite&context="+err.Error())
	}

	// TODO: check permissions
	_ = member

	code, err := logic.RandomString(32)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create an invite&context="+err.Error())
	}
	inviteCode := strings.ToLower("lki-" + logic.HashString(code + member.ID)[:16])

	invite, err := c.DB().GroupInvite.CreateOne(db.GroupInvite.ExpiresAt.Set(time.Now().Add(time.Hour*24*7)), db.GroupInvite.Group.Link(db.Group.ID.Equals(data.GroupID)), db.GroupInvite.CreatedBy.Link(db.User.ID.Equals(c.User().ID)), db.GroupInvite.Code.Set(inviteCode)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create an invite&context="+err.Error())
	}

	return c.Partial(http.StatusOK, group.InviteCard(invite))
}

type GroupInviteRedeemForm struct {
	InviteCode string `form:"invite-code"`
}

func RedeemGroupInvite(c *handlers.Context) error {
	var data GroupInviteRedeemForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to redeem an invite&context="+err.Error())
	}

	inviteCode := strings.ToLower(data.InviteCode)
	invite, err := c.DB().GroupInvite.FindFirst(db.GroupInvite.Code.Equals(inviteCode), db.GroupInvite.Enabled.Equals(true), db.GroupInvite.ExpiresAt.After(time.Now())).Exec(c.Request().Context())
	if err != nil {
		if db.IsErrNotFound(err) {
			log.Err(err).Str("inviteCode", inviteCode).Msg("failed to find an invite")
		}
		return c.Partial(http.StatusUnprocessableEntity, app.OnboardingGroups(map[string]string{"invite-code": data.InviteCode}, map[string]string{"invite-code": "invalid invite code"}))
	}
	if invite.UseCount >= invite.UseLimit {
		return c.Partial(http.StatusUnprocessableEntity, app.OnboardingGroups(map[string]string{"invite-code": data.InviteCode}, map[string]string{"invite-code": "invalid invite code"}))
	}

	existingMember, err := c.Member(invite.GroupID)
	if err != nil && !db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to redeem an invite&context="+err.Error())
	}
	if existingMember != nil {
		return c.Partial(http.StatusUnprocessableEntity, app.OnboardingGroups(map[string]string{"invite-code": data.InviteCode}, map[string]string{"invite-code": "you are already a part of this group"}))
	}

	member, err := c.DB().GroupMember.CreateOne(db.GroupMember.Group.Link(db.Group.ID.Equals(invite.GroupID)), db.GroupMember.User.Link(db.User.ID.Equals(c.User().ID)), db.GroupMember.Permissions.Set("v0/")).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to redeem an invite&context="+err.Error())
	}

	_, err = c.DB().GroupInvite.FindUnique(db.GroupInvite.ID.Equals(invite.ID)).Update(db.GroupInvite.UseCount.Increment(1), db.GroupInvite.RedeemedBy.Link(db.User.ID.Equals(c.User().ID))).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to redeem an invite&context="+err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/app/group/"+member.GroupID)
}
