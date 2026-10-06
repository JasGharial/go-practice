package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if(name == "") {
		return "", errors.New("Please provide a valid name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// RandomGreeting returns a random greeting message for the named person.
func RandomGreeting(name string) (string, error) {
	if(name == "") {
		return "", errors.New("Please provide a valid name")
	}

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	format := formats[rand.Intn(len(formats))]
	return fmt.Sprintf(format, name), nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := RandomGreeting(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}