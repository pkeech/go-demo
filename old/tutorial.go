// NAME PACKAGE
package tutorial

// IMPORT REQUIRED PACKAGES
import "fmt"

// FUNCTION: MAIN APPLICATION FUNCTION
func main() {
	// DEFINE VARIABLE
	var firstName, lastName string
	greeting := "Hello"
	first, second := 1, 2
	
	// DEFINE CONSTANTS (VARIABLES THAT CAN'T BE CHANGES)
	const PI = 3.14

	// SET VARIABLE
	firstName = "Peter"

	// OUTPUT VARIABLE
	fmt.Println(greeting)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(first)
	fmt.Println(second)
	fmt.Println(PI)

	// OUPUT HELLO WORLD 
	fmt.Println("Hello World")

	// CONDITIONAL (IF) STATEMENTS
	var num int
	fmt.Printf("Enter a Number: ")
	fmt.Scanln(&num)

	if num == 5 {
		fmt.Println("You Entered Number 5 !!")
	} else if num == 6 {
		fmt.Println("You Entered Number 6 !!")
	} else {
		fmt.Println("You DID NOT enter Number 5 or 6")
	}

}