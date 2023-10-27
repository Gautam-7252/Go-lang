package main

import "fmt"

type rectangle struct {
	Length  int
	Breadth int
}

func main() {
	Rect := rectangle{10, 20}
	Area := area(Rect.Length, Rect.Breadth)
	Perimeter := perimeter(Rect.Length, Rect.Breadth)
	fmt.Println("Area :", Area)
	fmt.Println("Perimeter :", Perimeter)
	Rect = rectangle{20, 40}
}
func area(len int, bread int) int {
	return len * bread
}
func perimeter(len int, bread int) int {
	return 2 * (len + bread)
}
