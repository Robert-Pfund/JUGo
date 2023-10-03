package files

import (
	"os"

	"github.com/Robert-Pfund/json-JUGo/utilities"
	"github.com/joho/godotenv"
	"github.com/tidwall/sjson"
)

// Konzept:
//

type Jug struct {
	data []byte
}

func SetupJSONFile() string {

	err := godotenv.Load()
	utilities.Check(err)

	var defaultfilename string = os.Getenv("DEFAULTFILENAME")

	var fileExists bool = checkForFile(defaultfilename)

	if !fileExists {
		jug, err := os.Create(defaultfilename)
		utilities.Check(err)
		jug.Close()
	}

	return defaultfilename
}

func checkForFile(filename string) bool {

	var exists bool

	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			exists = false
			return exists
		}
	}

	exists = true
	return exists
}

func Write(data []byte) {

	var location string = os.Getenv("DEFAULTFILENAME")

	file, err := os.Open(location)
	utilities.Check(err)
	defer file.Close()

	sjson.Set(string(data), location)

	os.WriteFile(location, data, 0644)
}

func ReadAll() {

	var location string = os.Getenv("DEFAULTFILENAME")

}
