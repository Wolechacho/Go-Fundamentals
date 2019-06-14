package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//we close the file, instead of forgetting later and waiting resources
	defer f.Close()
	r := strings.NewReader("Ali")
	//read the r value and write to the file
	io.Copy(f, r)

	//Open the file and read the values
	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
