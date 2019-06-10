package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

type fruit struct {
	name   string
	charac []string
}

//embedding a struct on another struct
type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type individual struct {
	first string
	last  string
	age   int
}

func (i individual) tell() {
	fmt.Println("I am, ", i.first, i.last, "and i am ", i.age, "years old")
}

// tell()

func main() {
	i1 := individual{
		first: "Steven",
		last:  "Mark",
		age:   30,
	}
	i1.tell()

	fr1 := fruit{
		name:   "Mango",
		charac: []string{"Sweet", "round", "red_colored"},
	}
	fr2 := fruit{
		name:   "Orange",
		charac: []string{"Juicy", "Fluid", "orange_colored"},
	}

	fmt.Println(fr1)

	for i, v := range fr1.charac {
		fmt.Printf("%v %v\n", i, v)
	}

	//storing the struct value into a map:
	m := map[string]fruit{
		fr1.name: fr1,
		fr2.name: fr2,
	}

	fmt.Println("this is the map: ", m)

	fmt.Println()
	fmt.Println()

	for _, v := range m {
		// fmt.Println(k)
		fmt.Printf("%v \n", v.name)
		for i, val := range v.charac {
			fmt.Printf("\t %v \t %v\n", i, val)
		}
	}
	fmt.Println()
	fmt.Println()

	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	fmt.Println(t)
	fmt.Println(t.doors)

	fmt.Println()
	fmt.Println()

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}
	fmt.Println(s)
	fmt.Println(s.doors)

	fmt.Println()
	fmt.Println()

	sal := secretAgent{
		person: person{
			first: "James",
			last:  "Magu",
			age:   20,
		},
		ltk: true,
	}
	p1 := person{
		first: "Steven",
		last:  "Ade",
		age:   100,
	}
	p2 := person{
		first: "Sam",
		last:  "Mark",
		age:   54,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first)
	fmt.Println(p2.last, p2.age)

	fmt.Println(sal)
	//we didnt have to do sal.person.first, but we can
	//the inner type get promoted to the outer type
	fmt.Println(sal.first, sal.age, sal.ltk)
	fmt.Println(p2.last, p2.age)

	//anonymous struct
	px := struct {
		first string
		last  string
		age   int
	}{
		first: "Goke",
		last:  "Phil",
		age:   10,
	}
	fmt.Println(px)

	fmt.Println()
	fmt.Println()

	//anonymous 2:
	anon := struct {
		t1 map[string]string
		t2 []string
	}{
		t1: map[string]string{
			"firstname": "Mike",
			"lastname":  "Ojo",
		},
		t2: []string{
			"Rice", "Beans", "Yam",
		},
	}

	fmt.Println(anon.t1)
	for i, v := range anon.t1 {
		fmt.Println(i, v)
	}

	fmt.Println()
	fmt.Println()

	bio := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Camsy": 12334443,
			"Obi":   39872667,
			"Romeo": 8877663,
		},
		favDrinks: []string{
			"pepsi",
			"Cocacola",
		},
	}
	fmt.Println(bio.first)
	fmt.Println(bio.friends)
	fmt.Println(bio.favDrinks)

	fmt.Println()

	for k, v := range bio.friends {
		fmt.Println(k, v)
	}
	for k, v := range bio.favDrinks {
		fmt.Println(k, v)
	}
}
