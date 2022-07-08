package main

import (
	"fmt"

	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Calesvol")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(message)

	names := []string{"Calesvol"}
	messages, errs := greetings.Hellos(names)

	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Println(messages)
}
