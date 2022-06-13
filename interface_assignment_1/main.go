package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	height float64
	base float64
}

type square struct {
	sideLength float64
}

func main(){
	a := triangle{height: 5, base: 10}
	b := square{sideLength: 10}
	printArea(a)
	printArea(b)
}

func (tri triangle) getArea() float64 {
	return 0.5 * tri.base * tri.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

