package str

import "strings"

func Cat(xs []string) string {
	s := xs[0]
	//we will clip the xs, ie: give us everything from position 1 onwards
	//including 1 itself
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

func Join(xs []string) string {
	return strings.Join(xs, " ")
}
