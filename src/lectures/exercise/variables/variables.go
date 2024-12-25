//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor string = "black"
	fmt.Println(favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	//var birthYear, age int = 1995, 30
	birthYear, age := 1995, 30
	fmt.Println("birthYear", birthYear, "age", age)

	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial string = "S"
		lastInitial  string = "M"
	)
	fmt.Println(firstInitial, lastInitial)

	var myAgeInDays int
	myAgeInDays = 10000
	fmt.Println("myAgeInDays", myAgeInDays)
}
