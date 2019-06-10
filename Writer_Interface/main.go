package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Fprintln(os.Stdout, "Hello sir")
	io.WriteString(os.Stdout, "Hello again")
}
