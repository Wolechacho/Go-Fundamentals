package saying

import (
	"fmt"
	"testing"
)

//T from package Testing
func TestGreet(t *testing.T) {
	s := Greet("James")
	if s != "Welcome my dear: James" {
		t.Error("got", s, "expected", "Welcome my dear: James")
	}
}

//Examples are tests
//example functions take no arguments and begin with the word Example instead of Test.
func ExampleGreet() {
	fmt.Println(Greet("James"))
}

//B from the package testing
func BenchmarkGreet(b *testing.B) {
	//code will be run many times
	//that variable N determines how many times it runs
	for i := 0; i < b.N; i++ {
		Greet("James")
	}
}
