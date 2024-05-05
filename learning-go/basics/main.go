package main

import "fmt"

func main() {
	var age int = 32
	fmt.Println(getA(&age))
	fmt.Println(age)
}

func getA(age *int) int {
	*age = *age - 18
	return *age
}
