package main

import "fmt"

// Google Doc: https://docs.google.com/document/d/1ckYpi6hcRkaBUEk975f54oGsHYHu7GhzOk7-nOrkNxo/edit

var x int = 4
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println("values: ", s)

	type t string

}
