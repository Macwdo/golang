package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHellos(t *testing.T) {
	names := []string{"Danilo"}
	want := regexp.MustCompile(`\b` + names[0] + `\b`)
	msg, err := Hello("Danilo")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Danilo") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
