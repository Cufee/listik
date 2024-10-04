package api

import (
	"fmt"
	"net/http"

	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

type CommonPathParams struct {
	ItemID  string `param:"itemId"`
	ListID  string `param:"listId"`
	GroupID string `param:"groupId"`
}

type ListCreateForm struct {
	Name        string `query:"name"`
	Description string `query:"description"`

	CommonPathParams
}

func CreateList(c *handlers.Context) error {
	var data ListCreateForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}
	// Check if a user belong to this group
	member, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	if data.Name == "" {
		data.Name, err = logic.RandomName()
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
		}
	}

	// Create a list
	list, err := c.DB().List.CreateOne(db.List.Name.Set(data.Name), db.List.Group.Link(db.Group.ID.Equals(member.GroupID)), db.List.CreatedBy.Link(db.User.ID.Equals(c.User().ID)), db.List.Desc.Set(data.Description)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/app/group/%s/list/%s", list.GroupID, list.ID))
}

func ListSetComplete(c *handlers.Context) error {
	var data CommonPathParams
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to update a list&context="+err.Error())
	}

	// Check if a user belong to this group
	_, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to update a list&context="+err.Error())
	}

	// Update the list
	newValue := c.QueryParam("checked") == "true"
	list, err := c.DB().List.FindUnique(db.List.ID.Equals(data.ListID)).With(db.List.Group.Fetch()).Update(db.List.Complete.Set(newValue)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	items, err := c.DB().ListItem.FindMany(db.ListItem.ListID.Equals(list.ID)).OrderBy(db.ListItem.Checked.Order(db.ASC), db.ListItem.Name.Order(db.ASC)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=group not found&context="+err.Error())
	}

	return c.Page(http.StatusOK, app.List{List: list, Group: list.Group(), Items: items, ShoppingMode: c.QueryParam("mode") != "edit"}.Render())
}
