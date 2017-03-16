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
import "testing"

/* the function to test in the string_utils package is named
   Reverse; our test will be referred to that function, checking
   the correct transformation on a given set of "in" strings (can
   be more than one in the same struct assignment. The expected result
   will be specified the same way, referencing the "want" string.
 */
func TestReverse(t *testing.T) {
	// definition of the struct containing in and want strings
	cases := []struct {
		in, want string
	}{
		// first check, giving the in, expecting the want
		{"This is a string", "gnirts a si sihT"},
		// in test want tset
		{"test", "tset"},
		// special case, when in empty, want should be the same
		{"", ""},
		// spacial case, same value for in and want
		{"anna", "anna"},
	}
	// now loop for each pair of values in the cases struct
	for _, c := range cases {
		/* got is assigned with the value of the tested
		   Reverse function on the "in" current value
		 */
		got := Reverse(c.in)
		/* checking if got hasn't the same vale of "want"
		   for the current c elementt of the struct
		  */
		if got != c.want {
			// In case of error, the exception is printed
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		} else {
			// Otherwse we also log the successful test
			t.Logf("Successful check on Reverse(%q), in was %q, want was %q",
			c.in, got, c.want)
		}
	}
}