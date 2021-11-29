package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//controls the program's state
	var running bool = true

	var userInput bool = true

	var showEvenness bool = true

	var showPrimeness bool = true

	goal := 16

	scanner := bufio.NewScanner(os.Stdin)

	for running {

		userInput = true

		//user input loop
		for userInput {
			fmt.Println("-----------------------------------------------------------------")
			if showEvenness || showPrimeness {
				fmt.Printf("Type a number to find it's")

				//this is structured like this because I am simply making a basic program to learn Go

				if showEvenness {
					fmt.Printf(" evenness")
				}

				if showEvenness && showPrimeness {
					fmt.Printf(" and")
				}

				if showPrimeness {
					fmt.Printf(" primeness")
				}

				fmt.Printf(". You can also type %q or %q to disable/enable their output. Type %q to process. Type %q to exit.\nPlease type in the highest number you want to solve to: ", "even", "prime", "solve", "exit")

				scanner.Scan()

			} else {
				fmt.Printf("You have disabled everything. Please re-enable one with %q or %q: ", "even", "prime")

				scanner.Scan()
			}

			input := scanner.Text()

			numberInput, _ := strconv.ParseInt(input, 10, 64)

			//ignore everything and exit the program
			if input == "exit" {
				os.Exit(1337)
			}

			if numberInput > 2 && (showPrimeness || showEvenness) {

				goal = int(numberInput)

				fmt.Printf("The number is now set to %q.\n", input)
			} else {

				if input == "even" {
					showEvenness = !showEvenness
				} else if input == "prime" {
					showPrimeness = !showPrimeness
				}

				//exits the user input loop if one of the settings are enabled
				if input == "solve" && (showPrimeness || showEvenness) {
					userInput = false
				}

			}
		}

		for x := 0; x <= goal; x++ {

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
