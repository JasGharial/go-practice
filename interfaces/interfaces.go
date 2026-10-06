package main

import (
	"log"
)

type I interface {
	M()
}

func main() {
	log.SetPrefix("Interfaces: ")
	log.SetFlags(0)
	log.Println("Package Invoked")

	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	log.Printf("(%v, %T)\n", i, i)
}
