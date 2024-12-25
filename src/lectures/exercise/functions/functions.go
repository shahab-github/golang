//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:

//* Call every function at least once

package main

import "fmt"

//   - Write a function that accepts a person's name as a function
//     parameter and displays a greeting to that person.
func greetPerson(name string) {
	fmt.Println("Hello Mrs", name)
}

//   - Write a function that returns any message, and call it from within
//     fmt.Println()
func returnSomething() string {
	return "Some good message"
}

//   - Write a function to add 3 numbers together, supplied as arguments, and
//     return the answer
func addThree(a, b, c int) int {
	return a + b + c
}

// * Write a function that returns any number
func five() int {
	return 5
}

// * Write a function that returns any two numbers
func twoTwos() (int, int) {
	return 2, 2
}

func main() {
	fmt.Println(returnSomething())

	greetPerson("John")

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	a, b := twoTwos()
	fivevalue := five()

	sum := addThree(a, b, fivevalue)
	fmt.Println(sum)

	newSum := addThree(five(), a, b)
	fmt.Println(newSum)

}
