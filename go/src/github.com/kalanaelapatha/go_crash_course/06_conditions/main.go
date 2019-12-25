package main

import "fmt"

func main() {

	x := 10
	y := 10

	if x <= y {
		fmt.Printf("%d is less than or equel %d\n", x, y)
	} else {

		fmt.Printf("%d is less than %d\n", y, x)
	}

	// Switch

	color := "yellow"

	switch color {

	case "red":
		fmt.Println("Colour is Red")

	case " blue":
		fmt.Println("Colour is Blue")

	default:
		fmt.Println("It's NOT blue or Red")
	}

}
