/*
   Library     : reverse_string.go
   Author      : Fabrizio Sgura fsgura@psl.com.co
   Description : This is part of an example library
                 called string_utils, including utilities
                 to perform operations on strings
   Site        : http://www.github.com/fsgura/testing-go
 */

/* Our package string_utils contains utility functions
   to perform custom operations with strings.
 */
package string_utils

/* The function Reverse returns its argument string reversed
   rune-wise left to right.
   i.e. string will be returned as gnirst
   The declaration means:
   function Reverse accepting an argument s of type string
   will return a string executing the inner code
 */
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
