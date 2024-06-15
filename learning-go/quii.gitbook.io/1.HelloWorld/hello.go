package main

import (
	"fmt"
)

// Defining a constant for the English greeting prefix
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello function that takes a name as a parameter and returns a greeting message
func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	return getLangPrefix(language) + name
}

func getLangPrefix(language string) string {
	prefix := englishHelloPrefix
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}
	return prefix
}

func main() {
	// Printing the greeting message for the name "krishna"
	fmt.Println(Hello("krishna", "English"))
}
