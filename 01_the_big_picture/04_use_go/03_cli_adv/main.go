// Go has pointers, but they are not as flexible and powerful(and also unsafe) as in C/C++
// coz there are a lot of limitations on pointers, and they are present in Go just for the sake of efficiency

// "02_cli" can only filter for "ERROR" values
// application will be more useful if we can customize the log file and the error level that we are interested in
// this customization is provided via command line parameters

// now that the application has command line parameters, to know how to use the application, you can do
// go run . -help

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// you don't need to pass arguments to "main" function to access command line paramters, "flag" package is enough
	path := flag.String("path", "myapp.log", "The path to the log that you want to analyse")
	// gets a pointer to command line parameter "path" whose default value is "myapp.log" and the help msg is "The path to the log that you want to analyse"
	level := flag.String("level", "ERROR", "Log level to search for: DEBUG, INFO, ERROR and CRITICAL")
	flag.Parse() // until this called, the variables "path" and "level" are not populated

	f, err := os.Open(*path) // need to dereference as "path" is a pointer
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s)
		}
	}
}
