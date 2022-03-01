// table-driven test is not a feature or a command or an api
// it just a pattern that is used to write testing code

package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	scenarios := []struct { // slice of a struct of the type having two string fields: "input" and "expect"
		input  string
		expect string
	}{ // defining the struct slice
		{
			input:  "Gopher",
			expect: "Hello, Gopher!\n",
		},
		{
			input:  "", // good to test with null inputs
			expect: "Hello, !\n",
		},
	}
	for _, s := range scenarios {
		got := Greet(s.input)
		if got != s.expect {
			t.Errorf("Did not get expected result for input %v. Expected %q, got: %q\n", s.input, s.expect, got)
		}
	}
	// an efficient way to test a specific api
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Bye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}
