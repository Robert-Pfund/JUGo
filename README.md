## JUGo

JUGo is a simple Database written in Go. It was created for small side-projects and is mainly focused on storing/deleting structs in/from a json-File.

### Installation

```sh
go get github.com/JUGo
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
	domain.Write("001", b1)

	// Read some data from json-File by id
	domain.Get("001")

	//--------------------------------------

	b2 := Booking{
		Firstname: "Donald",
		Lastname:  "Johnson",
	}

	domain.Write("002", b2)

	// Read all data from json-File
	domain.GetAll()

	//--------------------------------------

	domain.GetAll()

	// Delete some data from json-File by id
	domain.Delete("001")

	domain.GetAll()
}
```