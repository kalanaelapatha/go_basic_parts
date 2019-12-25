package main

import "fmt"

func main() {
	// //Arrays
	// var fruitArr [2]string

	// // Asssig Values

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	fruitArr := []string{"Apple", "Orange", "grape"}

	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[0:3])
}
