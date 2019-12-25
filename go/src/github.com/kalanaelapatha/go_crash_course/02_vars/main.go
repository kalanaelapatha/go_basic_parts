package main

import "fmt"

func main() {

	name, email := "kalana", "kalana@loops.lk" //outiside the function it will never works like that

	const isCool = true

	fmt.Println(name, email)
	fmt.Printf("%T\n", name)

}
