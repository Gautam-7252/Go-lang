//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

// Mathematical operations must be defined as constants using iota
const (
	Addition = iota
	Subtraction
	Multiply
	Divide
)

type operation int

func (op operation) calculate(A int, B int) {
	if op == Addition {
		fmt.Println("Addition of", A, B, "is :", A+B)
	} else if op == Subtraction {
		fmt.Println("Subtraction of", A, B, "is :", A-B)
	} else if op == Multiply {
		fmt.Println("Multiplication of", A, B, "is :", A*B)
	} else if op == Divide {
		fmt.Println("Division of", A, B, "is :", A/B)
	}
}

func main() {
	add := operation(Addition)
	add.calculate(50, 20)
	Subtract := operation(Subtraction)
	Subtract.calculate(50, 20)
	Multi := operation(Multiply)
	Multi.calculate(50, 20)
	Divide := operation(Divide)
	Divide.calculate(50, 20)
}
