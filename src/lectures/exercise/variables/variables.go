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

	//Requirements:
	//* Store your favorite color in a variable using the `var` keyword
	var favoriteColor string = "Black"
	fmt.Println(favoriteColor)
	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	year, age := 95, 28
	fmt.Println(year, age)
	//* Store your first & last initials in two variables using block assignment
	var (
		firstInitial = "S"
		lastInitial  = "M"
	)
	fmt.Println(firstInitial, lastInitial)

	//* Declare (but don't assign!) a variable for your age (in days),
	var ageInDays int
	//  then assign it on the next line by multiplying 365 with the age
	ageInDays = 365 * 29
	fmt.Println(ageInDays)
	// 	variable created earlier
	//
}
