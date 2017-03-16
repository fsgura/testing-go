package first_program

import "fmt"

func main() {
	fmt.Printf("This is a printed string, without end-of-line, using Printf. ")
	fmt.Printf("For that reason this other line is printed inline following the previous ! But ... \n")
	fmt.Println("At the end of the last line there was an explicit \\n, so that text was closed with a line-feed")
	fmt.Println("The previous text and this line is printed using Println, and this automatically enter the eol for us")
	fmt.Println(" ... and this was just a simple (maybe too much) go experience")
	fmt.Printf("Cheers and ... c u next time ;-) \n")
}
