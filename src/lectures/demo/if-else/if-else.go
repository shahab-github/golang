package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("Quiz 1 is greater than Quiz 2")
	} else if quiz1 < quiz2 {
		fmt.Println("Quiz 1 is less than Quiz 2")
	} else {
		fmt.Println("Quiz 1 nad Quiz 2 are equal")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println("acceptable grades")
	} else {
		fmt.Println("not acceptable grades")
	}
}
