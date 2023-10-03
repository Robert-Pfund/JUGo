package domain

import (
	"encoding/json"
	"log"
	"os"

	"github.com/Robert-Pfund/json-JUGo/utilities"
	"github.com/joho/godotenv"
	"github.com/tidwall/gjson"
)

type Jug struct {
	Items []Data
}

type Data struct {
	ID      string
	Jugdata JugData
}

type JugData interface {
}

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

func Write(data JugData) {

	var location string = os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	/*
		json, err := json.Marshal(data)
		utilities.Check(err)
	*/

	id := "001"

	datalist := make([]Data, 1)
	datalist = append(datalist, Data{
		ID:      id,
		Jugdata: data,
	})

	jug := &Jug{
		Items: datalist,
	}

	json, err := json.Marshal(jug)
	utilities.Check(err)

	log.Printf("JSON from Write: %s\n", data)

	//os.WriteFile(location, json, 0644)
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

// Check/Get Hash from json-File
/*
func BytesFromFile() string {

	var location string = os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	raw, err := os.ReadFile(location)
	utilities.Check(err)

	var bytes bytes.Buffer
	gob.NewEncoder(&bytes).Encode(raw)
	return base64.StdEncoding.EncodeToString(bytes.Bytes())
}
*/
