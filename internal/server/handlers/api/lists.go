package api

import (
	"fmt"
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

type ListCreateForm struct {
	Name        string `form:"name"`
	Description string `form:"description"`

	ListID  string `param:"listId"`
	GroupID string `param:"groupId"`
}

func CreateList(c *handlers.Context) error {
	var data ListCreateForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	if len(data.Name) < 1 || len(data.Name) > 80 {
		return c.Page(http.StatusUnprocessableEntity, app.CreateListDialog(data.GroupID, true, makeInputsMap(data), map[string]string{"name": "list name should be between 1 and 80 characters"}))
	}
	if len(data.Description) > 80 {
		return c.Page(http.StatusUnprocessableEntity, app.CreateListDialog(data.GroupID, true, makeInputsMap(data), map[string]string{"description": "list description is limited to 80 characters"}))
	}

	// Check if a user belong to this group
	member, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	// Create a list
	list, err := c.DB().List.CreateOne(db.List.Name.Set(data.Name), db.List.Group.Link(db.Group.ID.Equals(member.GroupID)), db.List.CreatedBy.Link(db.User.ID.Equals(c.User().ID)), db.List.Desc.Set(data.Description)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/app/group/%s/list/%s", list.GroupID, list.ID))
}
