// if you do:
// go run .
// this will build the executable in a temp location and run it from there, you won't see any executable created by it in your present working directory
// to just build a go application in your present working directory without running it, you can do:
// go build .
// this will generate an executable whose name is the name of the module in the "go.mod" file in the present working directory

// from go 1.13, modules are an efficient way to organize go project source code
// a module is a directory that contains "go.mod" file which contains module configuration
// the simplest of the "go.mod" file will have module name and version of the go language
// it is a good practice to use the github location of the module as the module name coz commands such as run, get will go to the link, download the code and do the action
// this way you can avoid navigating within your local machine

package main

func main() {

}
