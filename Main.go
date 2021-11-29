package main

import "fmt"

func main() {
	for x := 0; x <= 10; x++ {

		//check if divisibility
		var flag bool = false

		for i := 2; i < x; i++ {
			if x%i == 0 {
				flag = true
				break
			}
		}

		if !flag {
			fmt.Printf("%v is a prime number\n", x)
		} else {
			fmt.Printf("%v is not a prime number\n", x)
		}

		if x%2 == 0 {
			fmt.Printf("%v is even\n", x)
		} else {
			fmt.Printf("%v is odd\n", x)
		}

		fmt.Println("-----------------------------")
	}
}
