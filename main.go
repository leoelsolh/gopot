package main

import (
	"fmt"
	"log"

	"gopot/src"
)

var attemptlog *log.Logger

func main() {
	f, err := src.LoadLogFile("v1.log")
	if err != nil {
		log.Printf("Failed to open log file: %s", err)
		return
	}
	defer f.Close()

	attemptlog = log.New(f, "", 0)
	fmt.Println("Server started on :2222 ...")
	src.StartServer()
}
