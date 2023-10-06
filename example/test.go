package example

import (
	"log"

	"github.com/Robert-Pfund/json-JUGo/domain"
)

type Booking struct {
	Firstname string
	Lastname  string
}

type Whatever struct {
	Idk   int
	Hello string
}

func RunTest() {
	/*
		booking1 := Booking{
			Firstname: "Maria",
			Lastname:  "Berg",
		}

		booking2 := Booking{
			Firstname: "Peter",
			Lastname:  "Altmeier",
		}
	*/
	/*
		booking3 := Booking{
			Firstname: "Otto",
			Lastname:  "Müller",
		}

		oopsie := Whatever{
			Idk:   4,
			Hello: "Hello",
		}
	*/

	domain.Connect()
	/*
		domain.Write("001", booking1)
		domain.Write("002", booking2)
	*/

	//d1 := domain.GetAll()
	//log.Print("d1: ")
	//log.Println(d1)
	//log.Println(d1[0])

	/*
		domain.Write("003", booking3)
		domain.Read()
		domain.Write("004", oopsie)
		domain.Read()
	*/

	d2 := domain.Get("002")
	log.Print("d2: ")
	log.Println(d2)

	/*
		id2 := domain.GetById("999")
		fmt.Println(id2)
	*/

	/*
		content, err := json.Marshal(jug1.Content)
		utilities.Check(err)

		b1 := &Booking{
			Firstname: string(content),
			Lastname:  string(content),
		}

		fmt.Println(b1)

	*/
	// Ausgabe:
	// &{{"Firstname":"Peter","Lastname":"Altmeier"} {"Firstname":"Peter","Lastname":"Altmeier"}}
	// --> weiter mit gjson
}
