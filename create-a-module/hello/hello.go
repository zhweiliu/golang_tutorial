package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set up prefix string for standard logger.
	log.SetPrefix("greetings: ")
	// Set up flag to empty log date, time for standard logger.
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Eric", "Nova", "Adam"}

	// Request a greeting message.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it to the console and exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned map or messages to the console.
	for _, message := range messages {
		fmt.Println(message)
	}
}
