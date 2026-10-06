package main

import (
	"fmt"
	"rsc.io/quote"
	"example.com/greetings"
	"log"
	"example.com/basics"
	"example.com/pointers"
)

func main() {
	// Logs Setup
	log.SetPrefix("file.go: ")
	log.SetFlags(0)

	// Generic Message Prints
	fmt.Println("Hello, World!")
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())

	// Single Greeting
	message, err := greetings.RandomGreeting("Gladys")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	//  Multiple Greetings
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if(err != nil) {
		log.Fatal(err)
	}

	fmt.Println(messages)

	// Basics Package Invocation
	log.Println("<========= Invoking Basics Package: START =========>")
	basics.Basic()
	log.Println("<========= Invoking Basics Package: END =========>")

	// Pointers Package Invocation
	log.Println("<========= Invoking Pointers Package: START =========>")
	pointers.PointersInvocation()
	log.Println("<========= Invoking Pointers Package: END =========>")
}
