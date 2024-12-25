package main

import "fmt"

// double takes the inputs as an integer and double and return it
func double(x int) int {
	return x + x
}

// add takes two inputs and return the sum
func add(lhs, rhs int) int {
	return lhs + rhs
}

// greet, it takes not input and return no output
func greet() {
	fmt.Println("Hello from the greet function !")
}

func main() {

	//calling greet function
	greet()

	dozen := double(6)
	fmt.Println("Dozen", dozen)

	// bakerDozen
	bakerDozen := add(dozen, 1)
	fmt.Println("Baker Dozen", bakerDozen)

	// another bakers dozen
	anotherBakerDozen := add(double(6), 1)
	fmt.Println("Another Baker Dozen", anotherBakerDozen)
}
