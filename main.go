package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// Create a new function, the entry point
func main() {
	fmt.Println("Hello World. \n", rand.Intn(200))
	newVar := "text"
	fmt.Println(newVar)
	result, err := intDiv(2, 5)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("THis is the result %v", result)
}

func intDiv(no1, no2 int) (int, error) {
	var err error
	result := no1 + no2
	if result == 0 {
		err = errors.New("The result is zero")
		return 0, err
	}
	return result, err
}
