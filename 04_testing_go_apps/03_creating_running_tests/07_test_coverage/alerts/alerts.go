// To check the test coverage, you can do
// go test -cover

// To check the coverage and write the report to a file
// go test -coverprofile <cover_rpt>.out
// then to read/analyse the report:
//     to see which functions are covered on cmd line:
//     go tool cover -func <cover_rpt>.out
//     to see which lines of code are covered on web browser:
//     go tool cover -html <cover_rpt>.out

// to generate a report of how many times a line has executed:
// go test -coverprofile <count_rpt>.out -covermode count
// to see the output on web browser(output shown as "high coverage" and "low coverage" are on a relative scale)
// go tool cover -html <count_rpt>.out

// ".out" is the format recommended by go tooling
// as it is in ".gitignore" as well, you don't see those files in this folder, run all the commands(in each of the package folder) and see for yourself

package alerts

import "fmt"

func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

func depart(name string) string {
	return fmt.Sprintf("Bye, %v\n", name)
}
