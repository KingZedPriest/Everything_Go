package main

import (
	"errors"
	"fmt"
	//"math/rand"
)

// Create a new function, the entry point
func main() {
	//fmt.Println("Hello World. \n", rand.Intn(200))
	newVar := "text"
	fmt.Println(newVar)
	result, err := intDiv(-5, 5)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("This is the result %v", result)
	}
	dealingArr()
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

// This creates an array of 3, of the value type of int32,
// and has a default value of [0, 0, 0]

func dealingArr() {
	var arr []int32 = []int32{10, 20, 30}
	var arr2 []int32 = []int32{4, 5}
	arr = append(arr, arr2...)
	fmt.Printf("\n The length of the array is length %v, and capacity %v", len(arr), cap(arr))
}
