package jugo

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"

	"github.com/Robert-Pfund/jugo/utilities"
	"github.com/joho/godotenv"
)

type Jug struct {
	ID      string
	Content JugData
}

type JugData interface{}

// Connect sets up files needed before calling other JUGo-functions:
//
// (1) .env-File (creates one, if not-existent) and sets DEFAULTFILENAME for the jug.json-File
//
// (2) the data-directory and jug.json-File in which the Jugs will be stored
func Connect() {

	defaultLocation := "data/jug.json"

	exists := utilities.CheckForFile(".env")
	if !exists {
		jug, err := os.Create(".env")
		utilities.Check(err)

		env, err := godotenv.Unmarshal("DEFAULTFILENAME=" + defaultLocation)
		utilities.Check(err)
		err = godotenv.Write(env, ".env")
		utilities.Check(err)

		err = jug.Close()
		utilities.Check(err)
	}

	file := utilities.SetupJSONFile()
	log.Printf("Data Storage set up at: %s\n", file)
}

// Save writes the given data under the given id to the json-File
func Save(id string, data JugData) {

	Write(id, data, 0)
}

// Edit rewrites data correlating to the given id to the json-File
//
// Internally Edit is just Save preceded by a Delete of the previous entry correlating to the given id
func Edit(id string, data JugData) {

	Write(id, data, 1)
}

// Write performs the actual storing of data in json-Files
func Write(id string, data JugData, mode int) {

	if mode != 0 {
		Delete(id)
	}

	var DB []Jug

	location := os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	filedata, err := os.ReadFile(location)
	utilities.Check(err)

	dec := json.NewDecoder(strings.NewReader(string(filedata)))
	for {

		if err := dec.Decode(&DB); err == io.EOF {
			break
		} else if err != nil {
			utilities.Check(err)
		}
	}

	jug := Jug{
		ID:      id,
		Content: data,
	}

	DB = append(DB, jug)

	dbJson, err := json.Marshal(DB)
	utilities.Check(err)

	log.Printf("JSON-DB from Write: %s\n", DB)
	err = os.WriteFile(location, dbJson, 0644)
	utilities.Check(err)
}

// GetAll returns a list of all Jugs saved to the json-File
func GetAll() []Jug {

	location := os.Getenv("DEFAULTFILENAME")
	var DB []Jug

	file, err := os.Open(location)
	utilities.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	filedata, err := os.ReadFile(location)
	utilities.Check(err)

	log.Printf("Data from GetAll: %s\n", filedata)

	dec := json.NewDecoder(strings.NewReader(string(filedata)))
	for {

		if err := dec.Decode(&DB); err == io.EOF {
			break
		} else if err != nil {
			utilities.Check(err)
		}
	}

	return DB
}

// Get returns the Jug correlating to the given id
func Get(id string) JugData {

	location := os.Getenv("DEFAULTFILENAME")
	var DB []Jug

	file, err := os.Open(location)
	utilities.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	data, err := os.ReadFile(location)
	utilities.Check(err)

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

// Delete removes the Jug correlating to the given id from the json-File
func Delete(id string) {

	var oldDb []Jug
	var newDb []Jug

	location := os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}(file)

	filedata, err := os.ReadFile(location)
	utilities.Check(err)

	dec := json.NewDecoder(strings.NewReader(string(filedata)))
	for {

		if err := dec.Decode(&oldDb); err == io.EOF {
			break
		} else if err != nil {
			utilities.Check(err)
		}
	}

	for i, j := range oldDb {
		if j.ID != id {
			newDb = append(newDb, oldDb[i])
		}
	}

	dbJson, err := json.Marshal(newDb)
	utilities.Check(err)

	log.Printf("JSON-DB from Write: %s\n", newDb)
	err = os.WriteFile(location, dbJson, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
