package example

import "github.com/Robert-Pfund/JUGo/domain"

type Booking struct {
	Firstname string
	Lastname  string
}

type Whatever struct {
	Idk   int
	Hello string
}

func RunTest() {

	// Setup jug.json and .env
	domain.Connect()

	//--------------------------------------

	b1 := Booking{
		Firstname: "Max",
		Lastname:  "Mustermann",
	}

	// Save some custom struct to json-File
	domain.Save("001", b1)

	// Read some data from json-File by id
	domain.Get("001")

	//--------------------------------------

	b2 := Booking{
		Firstname: "Donald",
		Lastname:  "Johnson",
	}

	domain.Save("002", b2)

	// Read all data from json-File
	domain.GetAll()

	//--------------------------------------

	domain.GetAll()

	// Delete some data from json-File by id
	domain.Delete("001")

	domain.GetAll()

	//--------------------------------------

	b2_new := Booking{
		Firstname: "Erich",
		Lastname:  "Haupt",
	}

	domain.Get("002")

	domain.Edit("002", b2_new)

	domain.Get("002")
}
