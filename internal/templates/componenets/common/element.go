package common

import (
	"github.com/a-h/templ"
)

type element struct {
	attributes templ.Attributes
}

func (e *element) Set(attr string, value any) {
	if e.attributes == nil {
		e.attributes = make(templ.Attributes)
	}
	e.attributes[attr] = value
}

func (e *element) SetMultiple(attrs templ.Attributes) {
	if attrs == nil {
		return
	}

	if e.attributes == nil {
		e.attributes = make(templ.Attributes)
	}
	for key, value := range attrs {
		e.attributes[key] = value
	}
}

func (e *element) addClass(classes ...string) {
	for _, class := range classes {
		e.Set("class", getAttr[string](e.attributes, "class")+" "+class)
	}
}

func getAttr[T any](attributes templ.Attributes, key string) T {
	value, ok := attributes[key].(T)
	if !ok {
		var noop T
		return noop
	}
	return value
}
