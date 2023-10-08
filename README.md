## JUGo

JUGo is a simple Database written in Go. It was created for small side-projects and is mainly focused on storing/deleting structs in/from a json-File.

! This project is still under developement and new features like only accepting unique ids might be added and bring change to some functionality of the database !

### Installation

```sh
go get github.com/Robert-Pfund/JUGo
```

### Usage

```sh
package main

import "github.com/Robert-Pfund/json-JUGo/domain"

type Booking struct {
	Firstname string
	Lastname  string
}

func main() {

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
```