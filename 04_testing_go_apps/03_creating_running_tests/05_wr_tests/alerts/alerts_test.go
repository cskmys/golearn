package alerts_test // blackbox testing

import (
	"05_wr_tests/alerts" // for blackbox testing you'll need to import package just like you would import any other package
	"testing"
)

func TestGreet(t *testing.T) {
	got := alerts.Greet("Gopher") // alerts::Greet is a public function and hence can be accessed in another package
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}

/*
func TestDepart(t *testing.T) {
	got := alerts.depart("Gopher") // throws compilation error
	// "alerts::depart" is a private function and hence cannot be accessed in another package
	expect := "Bye, Gopher\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}
*/
