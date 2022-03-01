package messages

import "testing"

func TestGreet(t *testing.T) {
	t.Skip("coz we have table-driven test for 'Greet', are skipping simple test") // test will be skipped with the string output
	// using "testing.T::SkipNow" doesn't require you to provide a string argument
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestGreetTableDriven(t *testing.T) {
	t.Parallel() // adding this will cause the test to run parallely
	scenarios := []struct {
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
		} else {
			t.Logf("Got expected result for input %v. Expected %q, got: %q\n", s.input, s.expect, got)
		}
	}
}

func TestDepart(t *testing.T) {
	t.Parallel() // adding this will cause the test to run parallely
	got := depart("Gopher")
	expect := "Bye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}
