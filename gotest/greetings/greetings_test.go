package greetings

import (
	"testing"
	"regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    name := "guigui"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("guigui")
    if !want.MatchString(msg) || err != nil {
        t.Fatalf(`Hello("guigui") = %q, %v, want match for %#q, nil`, msg, err, want)
    }
}
