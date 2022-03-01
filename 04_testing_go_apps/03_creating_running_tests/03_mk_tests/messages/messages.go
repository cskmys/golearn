// In Go,
// a function name starting with an upper case letter will make it public and can be accessed by other packages that import it
// a function name starting with a lower case letter will make it private and cannot be accessed by other packages that import it
// private function can only be accessed by code within the same package

// a package can span across multiple go files, all go files that start with "package <package>" belong to package "<package>" and is normally put in the same folder
// a module can have multiple packages(folder) and it has a "go.mod" file at the root
// all unit test(white-box/blackbox) files of a package are put in the same folder that contains the package code
// for integration and end-to-end tests another directory and/or package is used to contain them

package messages

import "fmt"

func Greet(name string) string { // func name starts with upper case letter, hence it is public
	return fmt.Sprintf("Hello, %v\n", name)
}

func depart(name string) string { // func name starts with lower case letter, hence it is private
	return fmt.Sprintf("Bye, %v\n", name)
}
