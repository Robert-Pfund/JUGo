package domain

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Robert-Pfund/json-JUGo/utilities"
	"github.com/joho/godotenv"
)

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

	var DB []Jug

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

func GetAll() []Jug {

	var location string = os.Getenv("DEFAULTFILENAME")
	var DB []Jug

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	data, err := os.ReadFile(location)
	utilities.Check(err)

	log.Printf("Data from Get: %s\n", data)

	dec := json.NewDecoder(strings.NewReader(string(data)))
	for {

		if err := dec.Decode(&DB); err == io.EOF {
			break
		} else if err != nil {
			utilities.Check(err)
		}
	}

	return DB
}

func Get(id string) JugData {

	var location string = os.Getenv("DEFAULTFILENAME")
	var DB []Jug

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	data, err := os.ReadFile(location)
	utilities.Check(err)

	//log.Printf("Data from Get: %s\n", data)

	dec := json.NewDecoder(strings.NewReader(string(data)))

	for {
		if err := dec.Decode(&DB); err == io.EOF {
			break
		} else if err != nil {
			utilities.Check(err)
		}

		log.Printf("jug Content from Get: %s\n", DB)
	}

	for i, j := range DB {
		if j.ID == id {
			return DB[i].Content
		}
	}

	return DB
}
