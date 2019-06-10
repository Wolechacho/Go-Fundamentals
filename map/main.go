package main

import (
	"fmt"
	"log"
)

func main() {
	m := map[string]int{
		"James": 32,
		"Maggi": 90,
	}
	fmt.Println(m)
	fmt.Println(m["James"])

	v, ok := m["Barn"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Jam"]; ok {
		fmt.Println("this is the if: ", v)
	}
	log.Println("unknown key ", v)

	m["mike"] = 45

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "mike")
	fmt.Println(m)

	if v, ok := m["James"]; ok {
		fmt.Println("value: ", v)
		delete(m, "James")
	}
	fmt.Println(m)

	fmt.Println()
	fmt.Println()

	// //the map has a key of string and a string
	// mk := map[string]string{
	// 	"Steven": "Basket Ball",
	// 	"Kumi":   "Soccer",
	// }

	// for i, v := range mk {
	// 	fmt.Printf("%v likes %v\n", i, v)
	// }

	//the map has a key of string and a slice of string
	mk := map[string][]string{
		"Steven": []string{"Basket Ball", "Table Tennis", "Coding"},
		"Kumi":   []string{"Soccer", "VolleyBall"},
	}
	//adding to the map
	mk["Ojo"] = []string{"Ice cream", "Food", "Grain"}

	//delete from the map
	delete(mk, "Steven")

	for i, v := range mk {
		fmt.Printf("%v likes: \n", i)
		for j, value := range v {
			fmt.Printf("\t%v likes %v\n", j, value)
		}
	}

}
