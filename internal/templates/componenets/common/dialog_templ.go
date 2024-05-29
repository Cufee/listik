// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.696
package common

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "time"
import "fmt"

type Dialog struct {
	StartOpen bool

	ID           string
	Body         templ.Component
	Header       templ.Component
	ActionButton templ.Component
}

func (d *Dialog) GetID() string {
	if d.ID == "" {
		return fmt.Sprintf("dialog-%d", time.Now().Unix())
	}
	return d.ID
}

func (d *Dialog) ShowScript() templ.ComponentScript {
	return showDialogScript(d.GetID())
}

func showDialogScript(id string) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_showDialogScript_123d`,
		Function: `function __templ_showDialogScript_123d(id){const dialog = document.getElementById(id)
    if (!dialog) return
    dialog.open = true;
		dialog.scrollIntoViewIfNeeded(true)
    document.querySelector("input")?.focus()
}`,
		Call:       templ.SafeScript(`__templ_showDialogScript_123d`, id),
		CallInline: templ.SafeScriptInline(`__templ_showDialogScript_123d`, id),
	}
}

func (dialog Dialog) Render() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<dialog class=\"modal\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if dialog.StartOpen {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" open")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" id=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(dialog.GetID())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/templates/componenets/common/dialog.templ`, Line: 35, Col: 69}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"modal-box flex flex-col gap-1 bg-base-200 rounded-xl\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if dialog.Header != nil {
			templ_7745c5c3_Err = dialog.Header.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if dialog.Body != nil {
			templ_7745c5c3_Err = dialog.Body.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><form method=\"dialog\" class=\"modal-backdrop bg-black bg-opacity-50\"><button>close</button></form></dialog> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if dialog.ActionButton != nil {
			templ_7745c5c3_Err = dialog.ActionButton.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
