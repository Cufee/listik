package bulma

import (
	"strings"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

type Option func(*options)

type Options []Option

type options struct {
	classes []string
	nodes   []g.Node
}

func (o Options) fromNodes(nodes []g.Node) []g.Node {
	var opts options
	for _, apply := range o {
		apply(&opts)
	}

	nodes = append(nodes, opts.nodes...)
	if len(opts.classes) > 0 {
		nodes = append(nodes, h.Class(strings.Join(opts.classes, " ")))
	}
	return nodes
}

func nodeWithBaseClass(opts Options, children []g.Node, base func(...g.Node) g.Node, class string) g.Node {
	return base(opts.Class(class).fromNodes(children)...)
}

func None() Options {
	return Options{}
}

func Class(classes ...string) Options {
	return Options{func(opts *options) {
		opts.classes = append(opts.classes, classes...)
	}}
}

func ClassIf(class string, condition bool) Options {
	if !condition {
		return None()
	}
	return Options{func(opts *options) {
		opts.classes = append(opts.classes, class)
	}}
}

func With(node g.Node) Options {
	return Options{func(opts *options) {
		opts.nodes = append(opts.nodes, node)
	}}
}

func (o Options) Class(classes ...string) Options {
	return append(o, Class(classes...)...)
}

func (o Options) ClassIf(class string, condition bool) Options {
	return append(o, ClassIf(class, condition)...)
}

func (o Options) With(node g.Node) Options {
	return append(o, With(node)...)
}
