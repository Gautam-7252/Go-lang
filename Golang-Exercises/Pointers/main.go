// --Summary:
//
//	Create a program that can activate and deactivate security tags
//	on products.
//
// --Requirements:
// * Create a structure to store items and their security tag state
//   - Security tags have two states: active (true) and inactive (false)
//
// * Create functions to activate and deactivate security tags using pointers
// * Create a checkout() function which can deactivate all tags in a slice
// * Perform the following:
//   - Create at least 4 items, all with active security tags
//   - Store them in a slice or array
//   - Deactivate any one security tag in the array/slice
//   - Call the checkout() function to deactivate all tags
//   - Print out the array/slice after each change
package main

import "fmt"

// Security tags have two states: active (true) and inactive (false)
const (
	Active   = true
	Deactive = false
)

// Create a structure to store items and their security tag state
type Items struct {
	Item  string
	State bool
}

// Create functions to activate and deactivate security tags using pointers
func deactivate(Product *Items) {
	fmt.Println("Deactivating Security Tag")
	Product.State = Deactive
}

// Create functions to activate and deactivate security tags using pointers
func activate(Product *Items) {
	fmt.Println("Activating Security Tag")
	Product.State = Active
}

// Create a checkout() function which can deactivate all tags in a slice
func checkout(list []Items) {
	fmt.Println("Deactivating all the Security tags at the Checkout")
	for i := 0; i < len(list); i++ {
		list[i].State = Deactive
	}
}
func main() {
	// Create at least 4 items, all with active security tags
	// Store them in a slice or array
	Itemlist := []Items{
		{"Marshmello", true},
		{"Basketball", true},
		{"Football", true},
		{"Cricketball", true},
	}
	fmt.Println(Itemlist)
	// Deactivate any one security tag in the array/slice
	deactivate(&Itemlist[1])
	// Print out the array/slice after each change
	fmt.Println(Itemlist)
	// Call the checkout() function to deactivate all tags
	checkout(Itemlist)
	fmt.Println(Itemlist)
	activate(&Itemlist[0])
	fmt.Println(Itemlist)
}
