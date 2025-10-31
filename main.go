package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rules()
	choice()
}

func rules() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\nPlease select the difficulty level\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")
}

func choice() {
	fmt.Println("Please Select Your Difficulty Level")
	var chances int
	fmt.Scan(&chances)
	game(chances)
}

func game(chances int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var randomNum int = r1.Intn(100)
	var input int
	fmt.Println(randomNum)
	if chances == 1 {
		fmt.Println("\nYou have chosen Easy Difficulty")
		fmt.Println("\nPlease now enter the number to guess")
		for i := 0; i < 10; i++ {
			fmt.Scan(&input)
			if randomNum == input {
				fmt.Println("\nCongratulations you have guessed the correct number in", i+1, "tries!!!")
				return
			} else {
				fmt.Println("YOu have guessed incorrect number!!")
				if input < randomNum {
					fmt.Println("The number is greater than the number you have entered!!")

				} else {
					fmt.Println("You number is smaller than the number you have enetered!!")
				}
			}

		}
	} else if chances == 2 {
		fmt.Println("\nPlease now enter the number to guess")
		for i := 0; i < 5; i++ {
			fmt.Scan(&input)
			if randomNum == input {
				fmt.Println("\nCongratulations you have guessed the correct number in", i+1, "tries!!!")
				return
			} else {
				fmt.Println("YOu have guessed incorrect number!!")
				if input < randomNum {
					fmt.Println("The number is greater than the number you have entered!!")

				} else {
					fmt.Println("You number is smaller than the number you have enetered!!")
				}
			}
		}

	} else if chances == 3 {
		fmt.Println("\nPlease now enter the number to guess")
		for i := 0; i < 3; i++ {
			fmt.Scan(&input)
			if randomNum == input {
				fmt.Println("\nCongratulations you have guessed the correct number in", i+1, "tries!!!")
				return
			} else {
				fmt.Println("YOu have guessed incorrect number!!")
				if input < randomNum {
					fmt.Println("The number is greater than the number you have entered!!")

				} else {
					fmt.Println("You number is smaller than the number you have enetered!!")
				}
			}
		}
	}

}
