/*
   Program     : revert_string_program.go
   Author      : Fabrizio Sgura fsgura@psl.com.co
   Description : This go program uses the function Reverse
   		 importing the package string_utils
                 It prints a given hardcoded string with
                 its characters reverted.
   Site        : http://www.github.com/fsgura/testing-go
 */

/* This is the main program, declared as the main package
 */
package main

import (
	// package "fmt" containing the Printf function
	"fmt"

	/* our string_utils package previously created,
           containing the Reverse function.
           Please note that if you're using a different
           github account or a different folder structure,
           the whole path containing the library should be
           updated. This exception could be verified when
           forking a repo, changing names and references.
           Check your configuration details in that case.
         */
	"github.com/fsgura/testing-go/string_utils"
)

func main() {
	/* variable of type string, implicitly declared and assigned
	   inside the main function.
	 */
	var stringToBeReverted = "This string will be reverted !"

	/* variable containing the reverted string, set calling
	   the Reverse function of the string_utils package
	  */
	var revertedString = string_utils.Reverse(stringToBeReverted)

	// Printing both results using Println
	fmt.Println("The original string to be reverted is: " +
		     stringToBeReverted)
	fmt.Println("The reverted string using Reverse from string_utils is: " +
	             revertedString)
}