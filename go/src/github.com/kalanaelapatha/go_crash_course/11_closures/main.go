package main

import "fmt"

func adder() func(int) int { // in here adder function returns the another function

	sum := 0

	return func(x int) int { // this part will be returend to the main

		sum += x
		return sum
	} // end of the returned  part
}

func main() {

	sum := adder() // this the some kind assining a fuction to varible

	for i := 0; i < 10; i++ {

		fmt.Println(sum(i))
	}

}
