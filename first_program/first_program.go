/*
   Program     : first_program.go
   Author      : Fabrizio Sgura fsgura@psl.com.co
   Description : This is the first program of the
                 testing-go tutorial.
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
	/* the imported "ftm" package name is referenced
	    before every related function call.
	    Note that when splitting a string in multiple
	    lines, it's necessary to concatenate all the
	    substrings with the + concatenation operator
	 */
	fmt.Printf("This is a printed string, " +
		          "without end-of-line, using " +
		          "Printf. ")
	fmt.Printf("For that reason this other " +
			  "line is printed inline following " +
		          "the previous ! But ... \n")
	fmt.Println("At the end of the previous line there " +
		       "was an explicit \\n, so that text " +
		       "was closed with a line-feed")
	fmt.Println("The previous text and this line " +
		       "are printed using Println, and this " +
		       "automatically enter the eol for us")
	fmt.Println(" ... and this was just a simple " +
		       "(maybe too much) go experience")
	fmt.Printf("Cheers and ... GOing to c u " +
		          "next time ;-) \n")
}
