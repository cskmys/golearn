// to run this go to folder where this file and its ".mod" file exists and
// execute: go run .

// in Go:
// having an unused variable doesn't lead to a warning but an error
// not handling an error is going to throw a warning
// semicolons are not needed to end a statement

package main // word "package" indicates that we are creating an executable

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
) // to import packages into current project

func main() { // entry point
	f, err := os.Open("myapp.log") // "Open" function from "os" package is used to open files
	// it returns 2 values, "f" is for the file handle and "err" is for the optional error state

	defer f.Close() // when "defer" is used it performs the action that follows it after the function that it is called in exits(either naturally or by bailing out)
	// hence, at the very beginning you can make sure that you don't forget to close the file and no need to write it at the very end

	// there are no exceptions in Go, hence if you fail to open the file you don't get an exception but rather you get an error return value
	// so, you need to test if "err" is non-null value or not and if it's not, you'll need to handle it
	if err != nil { // "nil" means a null pointer value
		log.Fatal(err) // to bail out of the application
	}

	r := bufio.NewReader(f) // "bufio" package allows you to read an input stream, in this case it is filestream
	for {                   // "for" without any expression following it, is an infinite for loop
		s, err := r.ReadString('\n') // use the reader to read input stream until a newline character is encountered
		// "s" will have the string until newline character and "err" will have pointer to an error
		if err != nil {
			break // just break out of the infinite loop upon an error
		}
		if strings.Contains(s, "ERROR") { // if string "s" contains the string "ERROR"
			fmt.Println(s) // print the string "s" with string "ERROR" on console
		}
	}
}
