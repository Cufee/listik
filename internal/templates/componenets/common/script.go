package common

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/a-h/templ"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

type embeddedScript struct {
	script templ.ComponentScript
	params []any
}

func (s *embeddedScript) writeArgs(w io.Writer) error {
	paramsLen := len(s.params)
	for i, param := range s.params {
		paramEncodedBytes, err := json.Marshal(param)
		if err != nil {
			return err
		}
		if _, err = w.Write(paramEncodedBytes); err != nil {
			return err
		}
		if i+1 != paramsLen {
			if _, err = io.WriteString(w, ", "); err != nil {
				return err
			}
		}
	}
	return nil
}

var m = minify.New()

func MinifyScript(input string) string {
	r := bytes.NewReader([]byte(input))
	w := bytes.NewBuffer(nil)
	err := js.Minify(m, w, r, nil)
	if err == nil {
		return w.String()
	}
	return ""
}

func (s *embeddedScript) minify() {
	s.script.Function = MinifyScript(s.script.Function)
}

func Script(script templ.ComponentScript, params ...any) *embeddedScript {
	return &embeddedScript{script: script, params: params}
}

func (s *embeddedScript) Embed() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		if _, err = io.WriteString(w, `<script type="text/javascript">`+"\r\n"+s.script.Function+"\r\n"+s.script.Name+"("); err != nil {
			return err
		}
		if err := s.writeArgs(w); err != nil {
			return err
		}
		if _, err = io.WriteString(w, ")\r\n</script>"); err != nil {
			return err
		}
		return nil
	})
}

func (s *embeddedScript) EmbedMinified() templ.Component {
	s.minify()
	return s.Embed()
}
