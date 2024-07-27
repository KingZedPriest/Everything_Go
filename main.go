package main

import (
	"fmt"
	"math/rand"
)

// Create a new function, the entry point
func intDiv(no1, no2 int) int {
	result := no1 + no2
	return result
}

func main() {
	fmt.Println("Hello World. \n", rand.Intn(200))
	newVar := "text"
	fmt.Println(newVar)
	result := intDiv(2, 5)
	fmt.Printf("THis is the result %v", result)
}
