package api

import (
	"net/http"

	"github.com/cufee/shopping-list/internal/server/handlers"
	components "github.com/cufee/shopping-list/internal/templates/componenets"
	"github.com/cufee/shopping-list/internal/templates/pages/app"
	"github.com/cufee/shopping-list/prisma/db"
)

type ItemCreateForm struct {
	Name        string   `form:"name"`
	Tags        []string `form:"tags"`
	Description string   `form:"description"`

	ListID  string `param:"listId"`
	GroupID string `param:"groupId"`
}

func CreateItem(c *handlers.Context) error {
	var data ItemCreateForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	if len(data.Name) < 1 || len(data.Name) > 14 {
		return c.RenderPartial(app.CreateListDialog(data.GroupID, true, map[string]string{"name": data.Name, "description": data.Description}, map[string]string{"name": "Item name should be between 1 and 14 characters"}))
	}
	if len(data.Description) > 80 {
		return c.RenderPartial(app.CreateListDialog(data.GroupID, true, map[string]string{"name": data.Name, "description": data.Description}, map[string]string{"description": "Item description is limited to 80 characters"}))
	}

	// Check if a user belong to this group
	member, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item&context="+err.Error())
	}

	// TODO: Check permissions
	_ = member

	// Create a list
	item, err := c.DB().ListItem.CreateOne(db.ListItem.CreatedBy.Link(db.User.ID.Equals(c.User().ID)), db.ListItem.Name.Set(data.Name), db.ListItem.List.Link(db.List.ID.Equals(data.ListID)), db.ListItem.Desc.Set(data.Description)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item&context="+err.Error())
	}

	return c.RenderPartial(components.ListItem{Item: item, GroupID: data.GroupID}.Render())
}

type ItemSetCheckedData struct {
	ItemID  string `param:"itemId"`
	ListID  string `param:"listId"`
	GroupID string `param:"groupId"`
}

func ItemSetChecked(c *handlers.Context) error {
	var data ItemSetCheckedData
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to update an item&context="+err.Error())
	}

	newValue := c.QueryParam("checked") == "true"

	// Check if a user belong to this group
	member, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to update an item&context="+err.Error())
	}

	// TODO: Check permissions
	_ = member

	// Create a list
	updatedItem, err := c.DB().ListItem.FindUnique(db.ListItem.ID.Equals(data.ItemID)).Update(db.ListItem.Checked.Set(newValue)).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to update an item&context="+err.Error())
	}

	return c.RenderPartial(components.ListItem{Item: updatedItem, GroupID: data.GroupID}.Render())
}
