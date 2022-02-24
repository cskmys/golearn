package main

import "fmt"

func main() {
	var arr [3]int // reverse order of C/C++
	arr[0] = 1     // indexing same as C/C++
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])
	/*
		fmt.Println(arr[3]) // throws compilation error
		// due to array bounds checking at compile time
	*/

	arr2 := [3]int{1, 2, 3} // shorter form of initialization
	fmt.Println(arr2)
}
