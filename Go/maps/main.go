package main

import "fmt"

func main() {

	// type1
	// var colors map[string]string // its just a declaration it will be assigned to nil
	// colors = make(map[string]string)
	// or
	// colors = map[string]string{
	// 	"r": "r1",
	// }

	// type2
	// colors := make(map[string]string)
	// colors["x1"] = "y"

	// type3
	// colors := map[string]string{
	// 	"red": "#",
	// }

	// type4
	// colors := map[string]string{}
	// colors["x1"] = "y"

	// delete(colors, "x")

	colors := make(map[int]int, 2)
	colors[1] = 25
	colors[2] = 27
	// fmt.Println(colors)

	for key := range colors {
		fmt.Println("key:value", key, colors[key])
	}

	// slices in maps
	// mp1 := make(map[int][]int)
	// mp1[0] = []int{1, 2, 3}
	// fmt.Println(mp1)

	// mp2 := make(map[int]*[]int)
	// mp2[0] = &[]int{1, 2, 3}
	// fmt.Println(mp2)

	// for key := range mp2 {
	// 	fmt.Println(key)
	// }

}
