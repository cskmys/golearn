// in Go,
// in testing we have immediate failure and non-immediate failure
// immediate is one where you would declare the failure and stop the current test(just the current test not the whole set of tests/test suite) using one of the methods "FailNow", "Fatal", and "Fatalf"
// non-immediate is one where you would indicate the failure  using one of the methods "Fail", "Error", and "Errorf", but you continue to execute the current test

package messages // white-box testing

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestDepart(t *testing.T) {
	got := depart("Gopher")
	expect := "Bye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

func TestFailureTypes(t *testing.T) {
	t.Error("Err signals a failed test but doesn't stop the rest of the current test from executing")
	t.Fatal("As error didn't end the test it was possible to execute this, however fatal will fail the test and stop the execution.\nHence, you'll not see any messages in the following lines")
	t.Error("As fatal stopped the test, this message will never be seen on screen")
}
