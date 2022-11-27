package main

import (
	"log"
	"os"
	"time"
)

func main() {
	log.Print("Hello World")

	logger := log.New(os.Stdout, "", 0)

	logger.Print("Time=" + time.Now().Format(time.RFC3339) + " Hello World")
}
