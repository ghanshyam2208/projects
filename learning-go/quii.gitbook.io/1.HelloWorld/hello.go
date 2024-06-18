// package main

// import (
// 	"fmt"
// )

// // Defining a constant for the English greeting prefix
// const englishHelloPrefix = "Hello, "
// const spanishHelloPrefix = "Hola, "
// const frenchHelloPrefix = "Bonjour, "

// // Hello function that takes a name as a parameter and returns a greeting message
// func Hello(name string, language string) string {
// 	if name == "" {
// 		return englishHelloPrefix + "world"
// 	}
// 	return getLangPrefix(language) + name
// }

// func getLangPrefix(language string) string {
// 	prefix := englishHelloPrefix
// 	switch language {
// 	case "Spanish":
// 		prefix = spanishHelloPrefix
// 	case "French":
// 		prefix = frenchHelloPrefix
// 	}
// 	return prefix
// }

// func main() {
// 	// Printing the greeting message for the name "krishna"
// 	fmt.Println(Hello("krishna", "English"))

// 	a := 10
// 	var b, c = &a, &a
// 	fmt.Println(b, c) // 0xa0b0020 0xa0b0020
// 	fmt.Println(&b, &c) // 0xa09e138 0xa09e140
// }

package main

import "fmt"

func update(m map[string]string) {
	m["test"] = "test2"
}
func main() {
	m := make(map[string]string)
	m["test"] = "test1"
	update(m)
	fmt.Println(m)
}
