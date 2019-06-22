//Package word provide custom functions for working with strings
package word

import "strings"

//UseCount returns the number of times each word is used in a string
func UseCount(s string) map[string]int {
	//below is better than using strings.Split()
	xs := strings.Fields(s)
	//the key is a string and the value is int
	m := make(map[string]int)
	for _, v := range xs {
		//the key for the map is each word, initial value is 0
		m[v]++
	}
	return m
}

//Count returns the number of words in a string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
