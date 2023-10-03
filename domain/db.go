package domain

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Robert-Pfund/json-JUGo/utilities"
	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
)

var DB []Jug

type Jug struct {
	ID      string
	Content JugData
}

type JugData interface{}

func Connect() { //*Jug {

	defaultlocation := "data/jug.json"

	var envexists bool = utilities.CheckForFile(".env")
	if !envexists {
		jug, err := os.Create(".env")
		utilities.Check(err)

		env, err := godotenv.Unmarshal("DEFAULTFILENAME=" + defaultlocation)
		utilities.Check(err)
		err = godotenv.Write(env, ".env")
		utilities.Check(err)

		jug.Close()
	}

	file := utilities.SetupJSONFile()
	log.Printf("Data Storage set up at: %s\n", file)
}

func Write(id string, data JugData) {

	var location string = os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	jug := &Jug{
		ID:      id,
		Content: data,
	}

	DB = append(DB, *jug)

	json, err := json.Marshal(DB)
	utilities.Check(err)

	log.Printf("JSON-DB from Write: %s\n", DB)
	os.WriteFile(location, json, 0644)
}

func Read() {

	var location string = os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	json, err := os.ReadFile(location)
	utilities.Check(err)

	log.Printf("JSON from Read: %s\n", json)

	s := gjson.Get(string(json), "Data")
	log.Printf("GJSON from Read: %s\n", s)
}
