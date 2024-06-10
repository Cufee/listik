package common

import "strings"

type Style struct {
	classes []string
}

func (s *Style) WithClass(classes ...string) {
	s.classes = append(s.classes, classes...)

}

func (s *Style) Class(classes ...string) string {
	return strings.Join(s.classes, " ")
}
