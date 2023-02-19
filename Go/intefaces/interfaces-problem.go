package main

import "fmt"

type Triangle struct {
	Base   float64
	Height float64
}

type Square struct {
	Side float64
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.Base * t.Height
}

func (s Square) getArea() float64 {
	return s.Side * s.Side
}

type Shape interface {
	getArea() float64
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}

func main() {

	t := Triangle{1.1, 2.1}
	s := Square{Side: 1.1}

	printArea(t)
	printArea(s)

	fmt.Println(t, s)

}
