package main

import "github.com/Robert-Pfund/jugo"

type Booking struct {
	Firstname string
	Lastname  string
}

func RunExample() {

	// Setup jug.json and .env
	jugo.Connect()

	//--------------------------------------

	b1 := Booking{
		Firstname: "Max",
		Lastname:  "Mustermann",
	}

	// Save some custom struct to json-File
	jugo.Save("001", b1)

	// Read some data from json-File by id
	jugo.Get("001")

	//--------------------------------------

	b2 := Booking{
		Firstname: "Donald",
		Lastname:  "Johnson",
	}

	jugo.Save("002", b2)

	// Read all data from json-File
	jugo.GetAll()

	//--------------------------------------

	jugo.GetAll()

	// Delete some data from json-File by id
	jugo.Delete("001")

	jugo.GetAll()

	//--------------------------------------

	b2New := Booking{
		Firstname: "Erich",
		Lastname:  "Haupt",
	}

	jugo.Get("002")

	jugo.Edit("002", b2New)

	jugo.Get("002")
}
