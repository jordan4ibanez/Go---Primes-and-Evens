package main

import "fmt"

func main() {

	//controls the program's state
	var running bool = true

	var showEvenness bool = true

	var showPrimeness bool = true

	for running {

		for x := 0; x <= 10; x++ {

			if showPrimeness {

				//check divisibility
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
			}

			if showEvenness {
				if x%2 == 0 {
					fmt.Printf("%v is even\n", x)
				} else {
					fmt.Printf("%v is odd\n", x)
				}
			}

			fmt.Println("-----------------------------")
		}
	}
}
