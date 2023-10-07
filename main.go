package main

import (
	"github.com/Robert-Pfund/JUGo/example"
)

type TestData struct {
	Dataset []byte
}

func main() {

	example.RunTest()
	/*
		domain.Connect()


		domain.Write([]byte("Hello World"))
		domain.Write([]byte("How are you doing?"))
		domain.Read()
		log.Println(domain.BytesFromFile())

		domain.Write(TestData{
			Dataset: []byte("I'm doing great!"),
		})
		domain.Read()
		log.Println(domain.BytesFromFile())

		domain.Write([]byte("How are you doing?"))
		domain.Read()
		log.Println(domain.BytesFromFile())
	*/
}
