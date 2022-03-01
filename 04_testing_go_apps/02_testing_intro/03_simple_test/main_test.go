// in Go,
// to use standard library's testing suite, you'll need to name test files as "<filename>_test.go" and the test functions as "Test<function_name>"("<function_name>" should start with an uppercase letter)
// declare "package <package_name>" for white-box tests, i.e. the test has access to all the entities within the package "<package_name>"
// declare "package <package_name>_test" for blackbox tests, i.e. the test doesn't have access to all the entities within the package "<package_name>"
// it will need to import the "<package>" just like any other package and test what are all publicly exposed

// these aren't just naming conventions/suggestions, these are a pre-requisite for the Go tooling to identify and run
// for ex, "<filename>_test.go" tooling will understand not to include that file in production build

// to run, at the pwd of the go file, do:
// go test .

package main_test // blackbox testing

import "testing" // standard library's package for testing support

func TestAdittion(t *testing.T) { // "t *testing.T" is mandatory as it is needed to give access to the test runner and helps interact with it
	got := 2 + 2
	expected := 4
	if got != expected {
		t.Errorf("Did not get expected result. Got %v instead of %v", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 2 - 2
	expected := 1
	if got != expected {
		t.Errorf("Did not get expected result. Got %v instead of %v", got, expected)
	}
}
