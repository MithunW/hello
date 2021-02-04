package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("Greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	// Get a greeting message and print it.

	fmt.Println(message)
}
