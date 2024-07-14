// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.731
package app

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"
import "github.com/cufee/shopping-list/prisma/db"
import "github.com/cufee/shopping-list/internal/templates/componenets/common"
import "github.com/cufee/shopping-list/internal/templates/componenets/group"
import "github.com/cufee/shopping-list/internal/templates/componenets/tags"

type ManageGroup struct {
	Group    *db.GroupModel
	Lists    []db.ListModel
	ItemTags []db.ItemTagModel
	Members  []db.UserModel
	Invites  []db.GroupInviteModel
}

func (props ManageGroup) Render() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = common.PageHeader(common.BreadcrumbsTitle(
			[]common.BreadCrumb{
				{Label: "Groups", Href: "/app"},
				{Label: props.Group.Name, Href: fmt.Sprintf("/app/group/%s", props.Group.ID)},
				{Label: "Manage"},
			},
		), common.WithDescription(props.Group.Desc)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"divider\"></div><div class=\"flex flex-col gap-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ManageGroupTags(props.Group.ID, props.ItemTags).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col gap-2\"><span class=\"font-bold text-xl\">Current Members</span><div class=\"flex flex-row gap-2 flex-wrap\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, member := range props.Members {
			templ_7745c5c3_Err = group.MemberCard(&member, props.Group.OwnerID != member.ID).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div><div class=\"flex flex-col gap-2\"><span class=\"font-bold text-xl\">Active Invites</span><div class=\"flex flex-row gap-2 flex-wrap\" id=\"group-invites-conteainer\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(props.Invites) < 1 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<span class=\"bg-base-200 rounded-xl py-2 px-4 grow text-center\" id=\"group-no-invites-card\">there are no active invites</span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			for _, invite := range props.Invites {
				templ_7745c5c3_Err = group.InviteCard(&invite).Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"flex justify-center\"><button class=\"btn btn-link btn-lg\" hx-trigger=\"click\" hx-swap=\"beforeend\" hx-target=\"#group-invites-conteainer\" hx-post=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("/api/groups/%s/invites", props.Group.ID))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/pages/app/manage.templ`, Line: 59, Col: 68}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-on::after-request=\"if(event.detail.xhr.status.toString().startsWith(&#39;2&#39;))document.getElementById(&#39;group-no-invites-card&#39;)?.remove()\">create a new invite</button></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ManageGroupTags(groupID string, groupTags []db.ItemTagModel) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col gap-2\" id=\"manage-group-tags\"><span class=\"font-bold text-xl\">Item Tags</span><div class=\"flex flex-row gap-2 flex-wrap items-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, tag := range groupTags {
			templ_7745c5c3_Err = tags.ItemTag(groupID, tag).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = tags.CreateItemTagDialog("#manage-group-tags", groupID, false, nil, nil).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

type ManageList struct {
	Group *db.GroupModel
	List  *db.ListModel
	Items []db.ListItemModel
}

func (props ManageList) Render() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = common.PageHeader(common.BreadcrumbsTitle(
			[]common.BreadCrumb{
				{Label: "Groups", Href: "/app"},
				{Label: props.Group.Name, Href: fmt.Sprintf("/app/group/%s", props.Group.ID)},
				{Label: props.List.Name, Href: fmt.Sprintf("/app/group/%s/list/%s", props.Group.ID, props.List.ID)},
				{Label: "Manage"},
			},
		), common.WithDescription(props.List.Desc)).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"divider\"></div>Manage List")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
