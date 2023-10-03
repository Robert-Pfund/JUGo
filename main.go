package main

import (
	"log"

	"github.com/Robert-Pfund/json-JUGo/files"
)

func main() {

	file := files.SetupJSONFile()
	log.Printf("Data Storage set up at: %s\n", file)
}
