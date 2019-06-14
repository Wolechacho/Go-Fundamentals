package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
		//since we want a value of type error returned, we do:
		return []byte{}, errors.New(fmt.Sprintf("There was an error in toJSON %v", err))
	}
	return bs, nil
}
