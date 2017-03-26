/*
   Program     : first_program.go
   Author      : Fabrizio Sgura fsgura@psl.com.co
   Description : This is the first program of the
                 testing-go tutorial, an "hello word"
                 level-2 example :-)
                 It simply print some lines of text
                 on screen
   Site        : http://www.github.com/fsgura/testing-go
 */

/* Our first_program.go is a stand-alone file
   The package main effectively define this as the
   main application
 */
package main

/* we import the fmt go base package, which contains
   many function to format and parse IO. In this case
   we'll use Printf and Println only.
 */
import "fmt"

/* a function is declared with the "func" acronym,
   followed by the function name.
 */
func main() {
	/*
	   Defining a variable, type string, that will be
	   used with Printf function, passing it as a %s
	   placeholder.
	   When initialized without assignment, the default
	   value is automatically ""(empty string)
	 */
	var stringPlaceholder string
	/*
	   We will use this other variable, initializing it
	   directly with a value. Go will recognize the
	   correct type, letting us free of specifying the
	   type
	 */
	var otherPlaceholder = "other placeholder"
	/*
	   The next is a numeric variable, explicitly left
	   without declaring the type, that will be used in
	   the last sentence we will print on screen before
	   ending the program.
	 */
	var lastNumericPlaceholder = 2017
	/* The imported "ftm" package name is referenced
	   before every related function call.
	   Note that when splitting a string in multiple
	   lines, it's necessary to concatenate all the
	   substrings with the + concatenation operator,
	   exactly like other languages multilines code
	   writing.
	 */
	fmt.Printf("This is a printed string, " +
		          "without end-of-line, using " +
		          "Printf, passing the placeholder " +
		          "\"%s\". ", stringPlaceholder)
	fmt.Printf("For that reason this other " +
			  "line is printed inline following " +
		          "the previous, including the otherPlaceholder " +
		          "variable \"%s\" ! But ... \n", otherPlaceholder)
	fmt.Println("At the end of the previous line there " +
		       "was an explicit \\n, so that text " +
		       "was closed with a line-feed.")
	fmt.Println("The previous text and this line " +
		       "are printed using Println, and this " +
		       "automatically enter the eol for us")
	fmt.Println(" ... and this was just a simple " +
		       "(maybe too much) go experience")
	fmt.Printf("The current year is \"%d\" ! " +
		          "Cheers from go, your variable values \"%s\" and \"%s\" " +
		          "and ... GOing to c u next time ;-) \n",
		          lastNumericPlaceholder,
		          stringPlaceholder,
		          otherPlaceholder)
}
