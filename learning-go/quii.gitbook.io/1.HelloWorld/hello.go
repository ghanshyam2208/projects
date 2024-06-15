package main

import "fmt"

// Defining a constant for the English greeting prefix
const englishHelloPrefix = "Hello, "

// Hello function that takes a name as a parameter and returns a greeting message
func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + "world"
	}
	return englishHelloPrefix + name
}

func main() {
	// Printing the greeting message for the name "krishna"
	fmt.Println(Hello("krishna"))
}
