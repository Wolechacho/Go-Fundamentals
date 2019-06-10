package main

import "fmt"

func main() {
	// x := []int{2, 3, 4, 5, 60, 9}

	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// //colon allows us to slice the slice
	// fmt.Println(x[1:3])

	//consining the length of the slice, always use < len(x) or <= len(x)-1
	// for i := 0; i < len(x); i++ {
	// 	fmt.Println(i, x[i])
	// }
	// //or
	// for i := 0; i <= len(x)-1; i++ {
	// 	fmt.Println(i, x[i])
	// }
	// x = append(x, 30, 90, 100)
	// fmt.Println(x)

	// y := []int{12, 22, 33, 44}

	// x = append(x, y...)

	// fmt.Println(x)

	// x = append(x[:2], x[4:]...)
	// fmt.Println(x)

	//////////////////////////////////////////
	// x := make([]int, 10, 10)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))
	// x[0] = 32
	// x[9] = 409
	// x = append(x, 211)
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(cap(x))

	////////////////////////////////////////////
	// 	jb := []string{"James", "Mike", "Choco"}
	// 	fmt.Println(jb)
	// 	mp := []string{"Ade", "Goke", "Fur"}
	// 	fmt.Println(mp)

	// 	xp := [][]string{jb, mp}

	// 	fmt.Println(xp)
	// ///////////////////////////////////////////////
	// 	d := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 	fmt.Println(d[:4])
	// 	fmt.Println(d[4:])
	// 	fmt.Println(d[2:7])

	// x := []int{1, 2, 3, 5}
	// x = append(x, 54)
	// fmt.Println(x)
	// x = append(x, 20, 30, 40)
	// fmt.Println(x)
	// y := []int{11, 22, 33, 44}
	// x = append(x, y...)
	// fmt.Println(x)
	// x = append(x[4:], x[2:3]...)
	// fmt.Println(x)

	// x := [][]
	cp := []string{"Adam", "Ojo"}
	cq := []string{"Mali", "Cake"}

	cx := [][]string{cp, cq}
	fmt.Println(cx)

	for i, v := range cx {
		fmt.Println("record: ", i)
		for j, v := range v {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, v)
		}
	}
}
