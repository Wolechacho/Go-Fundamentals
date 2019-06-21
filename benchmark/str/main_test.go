package str

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	const s = "the string"
	xs := strings.Split(s, " ")
	theString := Cat(xs)
	if theString != "the string" {
		fmt.Println("Got ", s, "Expected", "the string")
	}
}

func TestJoin(t *testing.T) {
	const s = "the string"
	xs := strings.Split(s, " ")
	theString := Join(xs)
	if theString != "the string" {
		fmt.Println("Got ", s, "Expected", "the string")
	}
}

func ExampleCat() {
	s := "Hello me"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
}

func ExampleJoin() {
	s := "Hello me"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
}

const s = "This is the boy"

var xs []string

func BenchmarkCat(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}
func BenchmarkJoin(b *testing.B) {
	xs = strings.Split(s, " ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}
