// in Go,
// constants must be initialized during declaration itself
// if constant data type is not mentioned during declaration, it is inferred dynamically(not static, that means datatype can change) at compile time
package main

import (
	"fmt"
)

func main() {
	const pi = 3.14
	fmt.Println(pi)
	/*
		pi = 1.2 // throws compilation error
		// "pi" is a constant, hence you cannot assign to it
	*/

	/*
		t, _ := strconv.Atoi(time.Now().String())
		const trial = t + 1 // throws compilation error
		// "t" is not evaluated at compiler but dynamically at runtime
		// for a "const" type, the rhs expression should be evaluated at compile time
	*/

	const c = 3
	fmt.Println(c + 3)   // now "c" is interpreted as "int"
	fmt.Println(c + 1.2) // "c" is interpreted as "float"
	// we have 2 types of floats: "float32" and "float64", here don't know which type it is

	const d int = 4
	fmt.Println(d + 3) // now "d" is declared as "int"
	/*
		fmt.Println(d + 1.2) // throws compilation error
		// "d" is declared as "int", so there's no room for dynamic type inference
		// you cannot add "int" and "float", hence compilation error
	*/
	fmt.Println(float32(d) + 2.4) // now "d" is converted to "float32"
}
