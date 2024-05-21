package components

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Icon(size int) g.Node {
	return h.SVG(
		h.Width(fmt.Sprintf("%dpx", size)),
		h.Height(fmt.Sprintf("%dpx", size)),

		g.Attr("fill", "none"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("transform", "matrix(-1,0,0,1,0,0)"),
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Raw(`<path d="M8 18C19.9545 18 20.9173 7.82917 20.9935 2.99666C21.0023 2.44444 20.54 1.99901 19.9878 2.00915C3 2.32115 3 10.5568 3 18V22" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>`),
		g.Raw(`<path d="M3 18C3 18 3 12 11 11" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>`),
	)
}
