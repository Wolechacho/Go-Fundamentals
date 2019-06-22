package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("got", n, "expected", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	fmt.Println(YearsTwo(10))
}

func ExampleYears() {
	fmt.Println(Years(10))
}

func ExampleYearsTwo() {
	fmt.Println(Years(70))

}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
