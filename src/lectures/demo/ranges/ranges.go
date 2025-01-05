package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for i, elem := range slice {
		fmt.Println(i, elem)

		for _, ch := range elem {
			fmt.Printf("   %q\n", ch)
		}
	}
}
