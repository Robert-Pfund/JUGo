package example

import "github.com/Robert-Pfund/json-JUGo/domain"

type Booking struct {
	Firstname string
	Lastname  string
}

func RunTest() {

	booking1 := Booking{
		Firstname: "Maria",
		Lastname:  "Berg",
	}

	booking2 := Booking{
		Firstname: "Peter",
		Lastname:  "Altmeier",
	}

	domain.Connect()
	domain.Write(booking1)
	domain.Write(booking2)
	domain.Read()
}
