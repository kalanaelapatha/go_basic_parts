package main

import "fmt"

func main() {

	ids := []int{25, 75, 25, 50}

	//Loops through ids

	// for i, id := range ids {
	// 	fmt.Printf("%d -ID : %d\n", i, id)
	// }

	for id := range ids {
		fmt.Printf("ID : %d\n", id)
	}

	//Add all ids together
	sum := 0
	for _, id := range ids {

		sum += id
	}

	fmt.Println("Sum", sum)

	//range with Map -> string

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	for k, v := range emails {

		fmt.Printf("%s : %s\n", k, v)
	}

	//Get the Name

	for k := range emails {

		fmt.Println("Name :" + k)
	}

}
