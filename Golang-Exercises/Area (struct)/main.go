// --Summary:
//
//	Create a program to calculate the area and perimeter
//	of a rectangle.
//
// --Requirements:
// * Create a rectangle structure containing its coordinates
// * Using functions, calculate the area and perimeter of a rectangle,
//   - Print the results to the terminal
//   - The functions must use the rectangle structure as the function parameter
//   - After performing the above requirements, double the size
//     of the existing rectangle and repeat the calculations
//   - Print the new results to the terminal
//
// --Notes:
// * The area of a rectangle is length*width
// * The perimeter of a rectangle is the sum of the lengths of all sides
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
}
func area(len int, bread int) int {
	return len * bread
}
func perimeter(len int, bread int) int {
	return 2 * (len + bread)
}
