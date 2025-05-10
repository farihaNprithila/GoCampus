package secondclass

import "fmt"

// Declaring constants
const Pi = 3.1416        // Type inferred as float64
const DaysInWeek int = 7 // Explicit type

// Declaring multiple constants using parentheses
const (
	Zero = 0
	One  = 1
	Two  = 2
)

// Using constants in expressions
const (
	A = 5
	B = 10
)

func Variables() {
	//implicit
	//var x = "Hello From GoCampus"

	//implicit using short variable declaration operator
	// x:= "implicit using short variable declaration operator"

	//explicit
	//var x string = "Hello From GoCampus

	//short variable declaration operator cannot be used for explicit casting
	//x string:="Hello from explicit using short variable declaration operator" // giving a compilation error

	//multiple variables could be defined together
	//var x, y int

	//multiple variables can be defined and initialized together
	var x, y int = 1, 2

	//fmt.Println(x)
	fmt.Println(x, y)

	// Basic constant usage
	fmt.Println("Pi:", Pi)
	fmt.Println("Days in a week:", DaysInWeek)

	// Multiple constants usage
	fmt.Println("Zero:", Zero, "One:", One, "Two:", Two)

	// Constants in expressions
	fmt.Println("Sum of A and B:", A+B)

	// Using iota to create enumerated constants
	const (
		Monday = iota //untyped int
		Tuesday
		Wednesday
	)
	fmt.Println("Monday:", Monday, "Tuesday:", Tuesday, "Wednesday:", Wednesday)
}
