// in Go,
// there is no manual pointer arithmetic, hence "<pointer variable> + 1" will throw an error

package main

import "fmt"

func main() {
	var firstNamePtr *string // starting a type with "*" will indicate that it's a pointer of that datatype
	// as pointer "firstNamePtr" is uninitialised, it will be a null pointer
	fmt.Println(firstNamePtr) // "nil" which is null pointer in Go

	/*
		firstNamePtr = "Sid" // throws compilation error
		// string "Sid" is a value not an address and "firstNamePtr" is a pointer to string type
		// hence "firstNamePtr" cannot be assigned with a value
	*/
	/*
		// "*" preceding points variable indicates dereferencing
		// once you dereference a pointer variable, you can use it as a value variable
		// hence now it doesn't cause compilation error if you assign string "Sid" which is a value to "*firstNamePtr"
		*firstNamePtr = "Sid" // compiles but throws runtime error
		// "firstNamePtr" is "nil" aka a null pointer
		// dereferencing a null pointer is a violation of memory and hence code crashes
	*/

	var lastNamePtr *string = new(string)
	*lastNamePtr = "Chakka"
	fmt.Println(lastNamePtr)  // just printing pointer prints the address
	fmt.Println(*lastNamePtr) // printing dereference of pointer prints the value that the pointer is referring to

	middleName := "Subramanyam"
	fmt.Println(middleName)
	middleNamePtr := &middleName // assign a pointer to the data held by the variable
	fmt.Println(middleNamePtr)   // print pointer address
	fmt.Println(*middleNamePtr)  // print the data at this address
	*middleNamePtr = "Stark"     // dereference the pointer to mutate the data at the address that it is pointing to
	fmt.Println(middleName)      // now the variable bound to that address will show updated value
	fmt.Println(middleNamePtr)   // address of data is still the same
	fmt.Println(*middleNamePtr)  // dereferencing pointer will yield the data at the address
}
