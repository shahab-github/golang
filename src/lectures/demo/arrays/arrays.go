package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Printf("%s is cleaned\n", room.name)
		} else {
			fmt.Printf("%s is not cleaned\n", room.name)
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)
	fmt.Println("Performing cleaning")
	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliness(rooms)
}
