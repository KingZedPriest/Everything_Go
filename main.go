package main

import (
	"errors"
	"fmt"
	"strings"
	//"math/rand"
)

// Create a new function, the entry point
func main() {
	//fmt.Println("Hello World. \n", rand.Intn(200))
	//newVar := "text"
	//fmt.Println(newVar)
	//result, err := intDiv(-5, 5)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else {
	// 	fmt.Printf("This is the result %v", result)
	// }
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
	// var arr []int32 = []int32{10, 20, 30}
	// var arr2 []int32 = []int32{4, 5}
	// arr = append(arr, arr2...)
	// fmt.Printf("\n The length of the array is length %v, and capacity %v", len(arr), cap(arr))

	// var arr1 []int32 = make([]int32, 3, 10)
	// fmt.Printf("\n This is the length of the new array %v and the capacity %v", len(arr1), cap(arr1))

	var myMap map[string]uint8 = make(map[string]uint8)

	myMap = map[string]uint8{"helen": 23, "paul": 44}

	// var name, ok = myMap["json"]

	// if ok {
	// 	fmt.Printf("\n This is the value of json %v", name)
	// } else {
	// 	fmt.Printf("\n There is nothing like json in the map")
	// }

	// Looping over
	for name := range myMap {
		fmt.Printf("\n Name: %v", name)
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("\n i %v", i)
	// }

	var names = []string{"c", "h", "a", "r", "l", "e", "s"}

	//Create a string builder variable
	var stringBuilder strings.Builder

	//Loop through the string
	for i := range names {
		stringBuilder.WriteString(names[i])
	}

	//Create a new variable for that string
	var lastStr = stringBuilder.String()
	fmt.Printf("\n %v", lastStr)
}
