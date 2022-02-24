// in Go,
// if you don't use a module present in import block, you won't get a warning but rather an error
// whitespace has no effect on compilation(just like C/C++), but source code is formatted using tabs, not spaces
// semicolons are automatically inserted at the end of every line, hence pascal style of indentation of braces will throw errors and KnR style of braces must be used

package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
