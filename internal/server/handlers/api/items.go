package api

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/cufee/shopping-list/internal/logic"
	"github.com/cufee/shopping-list/internal/server/handlers"
	"github.com/cufee/shopping-list/internal/templates/componenets/common"
	"github.com/cufee/shopping-list/internal/templates/componenets/list"
	"github.com/cufee/shopping-list/prisma/db"
)

type ItemCreateForm struct {
	Tags []string `form:"tags"`

	Name        string `form:"name"`
	Price       string `form:"price"`
	Quantity    int    `form:"quantity"`
	Description string `form:"description"`

	CommonPathParams
}

func formFieldError(c *handlers.Context, data ItemCreateForm, field string, message string) templ.Component {
	containerSelector := c.QueryParam("container")

	// the input form can have some different state by now, it can be expanded and etc.
	// instead of replacing the whole thing, we only target the fields which have an error
	c.Response().Header().Set("HX-Reswap", "outerHTML")
	c.Response().Header().Set("HX-Retarget", "#create-list-item-form-"+field)
	c.Response().Header().Set("HX-Reselect", "#create-list-item-form-"+field)

	return list.NewListItem(data.GroupID, data.ListID, containerSelector, makeInputsMap(data), map[string]string{field: message})

}

func CreateItem(c *handlers.Context) error {
	var data ItemCreateForm
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new list&context="+err.Error())
	}

	if len(data.Name) < 1 || len(data.Name) > 80 {
		return c.Partial(http.StatusUnprocessableEntity, formFieldError(c, data, "name", logic.StringIfElse(len(data.Name) < 1, "name cannot be blank", "name is too long")))
	}

	if len(data.Description) > 80 {
		return c.Partial(http.StatusUnprocessableEntity, formFieldError(c, data, "description", "description is limited to 80 characters"))
	}
	if len(data.Price) > 80 {
		return c.Partial(http.StatusUnprocessableEntity, formFieldError(c, data, "price", "price is limited to 80 characters"))
	}
	if data.Quantity < 0 {
		return c.Partial(http.StatusUnprocessableEntity, formFieldError(c, data, "quantity", "quantity cannot be negative"))
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

	var optional []db.ListItemSetParam
	if data.Description != "" {
		optional = append(optional, db.ListItem.Desc.Set(data.Description))
	}
	if data.Price != "" {
		optional = append(optional, db.ListItem.Price.Set(data.Price))
	}
	if data.Quantity > 0 {
		optional = append(optional, db.ListItem.Quantity.Set(data.Quantity))
	}

	// Create a list
	item, err := c.DB().ListItem.CreateOne(
		db.ListItem.CreatedBy.Link(db.User.ID.Equals(c.User().ID)),
		db.ListItem.Name.Set(data.Name),
		db.ListItem.List.Link(db.List.ID.Equals(data.ListID)),
		optional...,
	).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item&context="+err.Error())
	}

	_, err = c.DB().List.FindUnique(db.List.ID.Equals(data.ListID)).Update(db.List.UpdatedAt.Set(time.Now())).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to create a new item&context="+err.Error())
	}

	return c.Partial(http.StatusCreated, list.ListItem{Item: item, GroupID: data.GroupID}.Render())
}

func DeleteItem(c *handlers.Context) error {
	var data CommonPathParams
	if err := c.Bind(&data); err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to delete an item&context="+err.Error())
	}

	// Check if a user belong to this group
	_, err := c.Member(data.GroupID)
	if db.IsErrNotFound(err) {
		return c.Redirect(http.StatusTemporaryRedirect, "/app")
	}
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to delete an item&context="+err.Error())
	}

	// TODO: Check permissions
	_, err = c.DB().ListItem.FindUnique(db.ListItem.ID.Equals(data.ItemID)).Delete().Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to delete an item&context="+err.Error())
	}

	_, err = c.DB().List.FindUnique(db.List.ID.Equals(data.ListID)).Update(db.List.UpdatedAt.Set(time.Now())).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to delete an item&context="+err.Error())
	}

	return c.Partial(http.StatusOK, common.Blank(""))

}

type ItemSetCheckedData struct {
	ItemID string `param:"itemId"`

	CommonPathParams
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

	_, err = c.DB().List.FindUnique(db.List.ID.Equals(data.ListID)).Update(db.List.UpdatedAt.Set(time.Now())).Exec(c.Request().Context())
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/error?message=failed to update and item&context="+err.Error())
	}

	return c.Partial(http.StatusOK, list.ListItem{Item: updatedItem, GroupID: data.GroupID, ShoppingMode: c.QueryParam("mode") == "shopping"}.Render())
}
