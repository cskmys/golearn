package main

import "fmt"

func main() {
	var i int // you can declare without initializing a variable
	i = 42    // initializing, if you don't do this then variable will take its default value, not a junk value like in C/C++
	fmt.Println(i)

	var f float32 = 3.14 // declaration and initialization on same line
	fmt.Println(f)

	firstName := "Sid" // ":=" forces compiler to infer the type, hence no need of "var" and type as shown above
	// most commonly used way to work with variables
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4) // built-in complex number type allows doing complex math which has made Go useful for complex mathematical simulations
	fmt.Println(c)

	r, im := real(c), imag(c) // multiple declaration and assignment in same line
	fmt.Println(r, im)
}
