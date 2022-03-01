package alerts_test // blackbox testing, and we can check only the public api and not the private one
// here nb public api = nb private api, hence 50% coverage

import (
	"07_test_coverage/alerts"
	"testing"
)

func TestGreet(t *testing.T) {
	got := alerts.Greet("Gopher")
	expect := "Hello, Gopher!\n"
	if got != expect {
		t.Errorf("Did not get expected result. Wanted %q, got: %q\n", expect, got)
	}
}
