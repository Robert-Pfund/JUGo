package example

type Booking struct {
	Firstname string
	Lastname  string
}

type Whatever struct {
	Idk   int
	Hello string
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

	booking3 := Booking{
		Firstname: "Otto",
		Lastname:  "Müller",
	}

}

/*
		booking1 := Booking{
			Firstname: "Maria",
			Lastname:  "Berg",
		}

			booking2 := Booking{
				Firstname: "Peter",
				Lastname:  "Altmeier",
			}

			booking3 := Booking{
				Firstname: "Otto",
				Lastname:  "Müller",
			}

			oopsie := Whatever{
				Idk:   4,
				Hello: "Hello",
			}

	domain.Connect()

	domain.GetAll()

	domain.Delete("001")

	domain.GetAll()


	content, err := json.Marshal(jug1.Content)
	utilities.Check(err)

	b1 := &Booking{
		Firstname: string(content),
		Lastname:  string(content),
	}

	fmt.Println(b1)

// Ausgabe:
// &{{"Firstname":"Peter","Lastname":"Altmeier"} {"Firstname":"Peter","Lastname":"Altmeier"}}
// --> weiter mit gjson
*/
