package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 3}
		ella = Passenger{Name: "Ela", TicketNumber: 2}
	)
	fmt.Println(bill, ella)

	var hiedi Passenger
	hiedi.Name = "Hiedi"
	hiedi.TicketNumber = 4
	fmt.Println(hiedi)

	casey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has Boarded the bus")
	}
	if casey.Boarded {
		fmt.Println("Casey has Boarded the bus")
	}

	hiedi.Boarded = true
	bus := Bus{hiedi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
