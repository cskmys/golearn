// in Go,
// an array is of fixed size and nb elements can never be changed but a slice is of variable size

package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	// slices are used to refer to a part of an array(like references in C++)
	// any change in array is going to reflect in slice and the vice versa
	slice := arr[:] // slice "slice" refers to "[:]" i.e. start to end of "arr"
	fmt.Println(arr, slice)

	arr[1] = 42
	slice[2] = 27
	fmt.Println(arr, slice)

	// if you don't specify the size, compiler will infer it as a slice rather than an array
	slice2 := []int{1, 2, 3} // declaring a slice without the underlying array will cause the compiler to
	// create the underlying array, populate it with values in braces and create a slice and bind it to a variable name
	fmt.Println(slice2)

	slice2 = append(slice2, 4) // as slice is not a fixed size entity like an array, it can be extended
	// during append operation the underlying array is copied to new bigger array with new elements and the slice is updated to point to new underlying array
	// there are APIs that allow you to improve efficiency by preallocating space in the underlying array for a given slice which will be very useful of arrays of very very very big size
	// however for most applications, you can just let Go manage it
	fmt.Println(slice2)
	slice2 = append(slice2, 42, 27) // multiple elements can be added
	fmt.Println(slice2)

	// you can create a slice of another slice
	slice3 := slice2[1:]  // "[1:]" indicates a slice from the index 1 to end
	slice4 := slice2[:2]  // "[:2]" indicates a slice from the start to index 1(decrement the number after the colon)
	slice5 := slice2[1:2] // "[1:2]" indicates a slice from index 1 to index 1
	fmt.Println(slice3, slice4, slice5)
}
