package api

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/componenets/tags"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

type TagCreateOptions struct {
	Name        string `form:"name"`
	Description string `form:"description"`

	CommonPathParams
}

func CreateItemTag(c *handlers.Context) error {
	var data TagCreateOptions
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item tag&context="+err.Error())
	}

	if len(data.Name) < 1 || len(data.Name) > 20 {
		return c.Partial(http.StatusUnprocessableEntity, tags.CreateItemTagDialog("#manage-group-tags", data.GroupID, true, makeInputsMap(data), map[string]string{"name": "item tag names should be between 1 and 20 characters"}))
	}
	if len(data.Description) > 80 {
		return c.Partial(http.StatusUnprocessableEntity, tags.CreateItemTagDialog("#manage-group-tags", data.GroupID, true, makeInputsMap(data), map[string]string{"description": "item tag description is limited to 80 characters"}))
	}

	// Check if a user belong to this group
	member, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item tag&context="+err.Error())
	}

	// Create a tag
	_, err = c.DB().ItemTag.CreateOne(db.ItemTag.Name.Set(data.Name), db.ItemTag.CreatedBy.Link(db.User.ID.Equals(c.User().ID)), db.ItemTag.Group.Link(db.Group.ID.Equals(member.GroupID))).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item tag&context="+err.Error())
	}

	// Tags
	tags, err := c.DB().ItemTag.FindMany(db.ItemTag.GroupID.Equals(member.GroupID)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item tag&context="+err.Error())
	}

	return c.Partial(http.StatusCreated, app.ManageGroupTags(data.GroupID, tags))
}
