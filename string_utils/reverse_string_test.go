/*
   UnitTest    : test_reverse_string.go
   Author      : Fabrizio Sgura fsgura@psl.com.co
   Description : This is a go unit test for the library
                 reverse_string.go
   Site        : http://www.github.com/fsgura/testing-go
 */

/* The current unit test is still part of the string_utils,
   although the go build and install steps will ignore it
 */
package string_utils

/* importing the go base package testing, containing the
   unit test functions and utilities
 */
import (
	"testing"
	"fmt"
	"github.com/mgutz/ansi"
)

/* the function to test in the string_utils package is named
   Reverse; our test will be referred to that function, checking
   the correct transformation on a given set of "in" strings (can
   be more than one in the same struct assignment. The expected result
   will be specified the same way, referencing the "want" string.
 */
func TestReverse(t *testing.T) {
	success := ansi.Color("SUCCESS","green")
	failure := ansi.Color("FAILURE", "red+b")
	// definition of the go slice containing in and want strings
	cases := []struct {
		in, want string
	}{
		// first check, giving the in, expecting the want
		{"This is a string", "gnirts a si sihT"},
		// in test want tset
		{"test", "tset"},
		// special case, when in empty, want should be the same
		{"", ""},
		// special case, same value for in and want
		{"anna", "anna"},
		// more cases
		{"check this value", "eulav siht kcehc"},
	}
	var casesSliceCapacity = cap(cases)
	var casesSliceLength = len(cases)
	fmt.Printf("The current slice for \"cases\" has %d elements capacity\n", casesSliceCapacity)
	fmt.Printf("Going to perform %d tests...\n", casesSliceLength)
	var successCounter,failureCounter int
	successCounter = 0
	failureCounter = 0
	// now loop for each pair of values in the cases slice
	for _, c := range cases {
		/* got is assigned with the value of the tested
		   Reverse function on the "in" current value
		 */
		got := Reverse(c.in)
		/* checking if got hasn't the same value of "want"
		   for the current c element of the struct
		  */
		if got != c.want {
			// In case of error, the exception is printed
			t.Errorf("%s: Reverse(%q) == %q, want %q", failure, c.in, got, c.want)
			failureCounter++
		} else {
			// Otherwse we also log the successful test
			t.Logf("%s: check on Reverse(%q):\n\t\tcurrent value  : %q\n\t\texpected value : %q",
				success, c.in, got, c.want)
			successCounter++
		}
	}
	fmt.Printf("Total number of %s : %d\n", success, successCounter)
	fmt.Printf("Total number of %s : %d\n", failure, failureCounter)
}