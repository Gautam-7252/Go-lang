// --Summary:
//
//	Create a program that can store a shopping list and print
//	out information about the list.
//
// --Requirements:
//   - Using an array, create a shopping list with enough room
//     for 4 products
//   - Products must include the price and the name
//   - Insert 3 products into the array
//   - Print to the terminal:
//   - The last item on the list
//   - The total number of items
//   - The total cost of the items
//   - Add a fourth product to the list and print out the
//     information again
package main

import "fmt"

type info struct {
	//Products must include the price and the name
	Name  string
	Price int
}

func store(list [4]info) {
	var totalprice, count int
	fmt.Println("Info")
	for i := 0; i < len(list); i++ {
		totalprice += list[i].Price
		//Print to the terminal:
		if list[i].Name != "" {
			count += 1
		}
	}
	//The last item on the list
	fmt.Println("Last item :", list[count-1])
	//The total number of items
	fmt.Println("Total items :", count)
	//The total cost of the items
	fmt.Println("Total cost :", totalprice)
}
func main() {
	//Using an array, create a shopping list with enough room
	//for 4 products
	List := [4]info{
		//Insert 3 products into the array
		{"Glasses1", 100},
		{"Glasses2", 200},
		{"Glasses3", 300},
	}
	store(List)

	//Add a fourth product to the list and print out the
	//information again
	List[3] = info{"Glasses4", 400}
	store(List)
}
