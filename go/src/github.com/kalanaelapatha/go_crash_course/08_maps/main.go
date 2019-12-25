package main

import "fmt"

func main() {
	// //define maps
	// emails := make(map[string]string)

	// //Assign kv

	// emails["Boz"] = "nnn@gmail.com"
	// emails["chamika"] = "chamika@gmail.com"
	// emails["binura"] = "binura@gmail.com"

	//declare the map and kv

	emails := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	emails["chamika"] = "chamika@gmail.com" //after also can add values

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Boz"])

	//delete the map

	delete(emails, "Boz")
	fmt.Println(emails)

}
