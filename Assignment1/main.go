package main

import "fmt"

type square struct {
	sideLength float64
}
type triangel struct {
	height float64
	base float64
}

type shape interface {
	getArea() float64
}

func main() {
s := square{
	sideLength: 5,
}
t := triangel{
	height: 10,
	base: 10,
}
printArea(s)
printArea(t)
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

func (t triangel) getArea() float64 {
		return 0.5*t.base*t.height
}

func (s square) getArea() float64 {
		return s.sideLength*s.sideLength
}
