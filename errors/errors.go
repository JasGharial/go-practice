package main

import (
	"fmt"
	"time"
)

// Types
type MyError struct {
	When time.Time
	What string
}

// Receiver Methods
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// Main Functions
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
