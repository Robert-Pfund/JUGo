package utilities

import (
	"os"

	"github.com/joho/godotenv"
)

func SetupJSONFile() string {

	err := godotenv.Load()
	Check(err)

	defaultFilename := os.Getenv("DEFAULTFILENAME")

	fileExists := CheckForFile(defaultFilename)

	if !fileExists {
		jug, err := os.Create(defaultFilename)
		Check(err)
		err = jug.Close()
		Check(err)
	}

	return defaultFilename
}

func CheckForFile(filename string) bool {

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
