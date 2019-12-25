package main

import "fmt"

func main() {

	a := 5
	b := &a // point the a's memory address not the value of "a"
	fmt.Println(a, b)

	fmt.Printf("%T\n", a)

	fmt.Printf("%T\n", b)

	// to read the values of the pointers values

	fmt.Println(*b) //same value will come from both
	fmt.Println(*&a)

	//change the value of the  pointer

	*b = 10
	fmt.Printf("After change the values of the memory address  :")
	fmt.Println(a)
}
