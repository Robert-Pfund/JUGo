package utilities

import (
	"os"

	"github.com/joho/godotenv"
)

func SetupJSONFile() string {

	err := godotenv.Load()
	Check(err)

	var defaultfilename string = os.Getenv("DEFAULTFILENAME")

	var fileExists bool = CheckForFile(defaultfilename)

	if !fileExists {
		jug, err := os.Create(defaultfilename)
		Check(err)
		jug.Close()
	}

	return defaultfilename
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
