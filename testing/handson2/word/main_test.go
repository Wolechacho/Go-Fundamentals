package word

import (
	"fmt"
	"testing"

	"../quote"
)

func TestCount(t *testing.T) {
	n := Count("one two three")
	if n != 3 {
		t.Error("got", n, "Want: ", 3)
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three")
	for k, v := range m {
		//where k is the key, which represents the individual words
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 2 {
				t.Error("got", v, "want", 2)
			}
		}
	}
}

func ExampleCount() {
	fmt.Println(Count("one two three"))
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//because we are the word package that is why we didnt use word.Count
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
