package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "Steven",
		Last:    "Victor",
		Sayings: []string{"Nice", "Awesome", "Lovely"},
	}

	//JSon stringifying
	bs, err := json.Marshal(p1)
	if err != nil {
		// fmt.Println("this is the error: ", err)
		// return
		log.Fatalln("Json did not marshal - here's the error: ", err)
	}
	fmt.Println(string(bs))
}
