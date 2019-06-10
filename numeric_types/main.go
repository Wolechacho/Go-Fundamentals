package main

import (
	"fmt"
	"runtime"
)

var a int8 = 127
var b int8 = -128

func main() {

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
