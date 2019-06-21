//Package acdc asks if you are ready to rock
package acdc

//Sum is the function that have the  unimited value of type int and return the int:
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
