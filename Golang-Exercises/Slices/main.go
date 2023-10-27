//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

//Using a slice, create an assembly line that contains type Part
type part string

//Create a function to print out the contents of the assembly line
func show(line []part) {
	for i := 0; i < len(line); i++ {
		fmt.Println(line[i])
	}
}
func main() {
	//Create an assembly line having any three parts
	assemblyLine := []part{"one", "two", "three"}
	fmt.Println("3 parts :")
	show(assemblyLine)
	//Add two new parts to the line
	assemblyLine = append(assemblyLine, "four", "five")
	fmt.Println("5 parts :")
	show(assemblyLine)
	//Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("2 parts :")
	show(assemblyLine)
}
