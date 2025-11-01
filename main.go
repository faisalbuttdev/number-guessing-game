//Number Guessing Game
//I have tried to make this as functional as possbile

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	easylevel   = 1
	mediumlevel = 2
	hardlevel   = 3
	minNum      = 1
	maxNum      = 100
)

type difficulty struct {
	Name    string
	Chances int
}

var selectedDifficulty = map[int]difficulty{
	easylevel:   {"Easy", 10},
	mediumlevel: {"Medium", 5},
	hardlevel:   {"Hard", 3},
}

func main() {
	for {

		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomNum := randGen(r1)
		rules()
		chances := choice()
		game(chances, randomNum)
		if !playagain() {
			break
		}
	}
}

func rules() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\nPlease select the difficulty level\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")
}

func randGen(r1 *rand.Rand) int {
	var randomNum int = r1.Intn(maxNum) + minNum
	return randomNum
}

func choice() int {
	for {
		fmt.Println("Please Select Your Difficulty Level")
		var chances int
		_, err := fmt.Scan(&chances)
		if err != nil {
			fmt.Println("Please Enter a Number (1 , 2 or 3)")
			continue
		}

		if chances < 1 || chances > 3 {
			fmt.Println("Please Enter a Number Between 1-3")
			continue
		}
		return chances
	}
}

func game(chances, randomNum int) { //Control Structure for Choices

	diff, exists := selectedDifficulty[chances]
	if !exists {
		fmt.Println("Invalid Difficulty Selected")
		return
	}
	fmt.Println("So you have chosen the difficulty level ", diff.Name)
	gameplay_loop(randomNum, diff.Chances)
}

func gameplay_loop(randomNum, counter int) { //Handles Main Game Logic
	fmt.Println("\nPlease now enter the number to guess")
	var input int
	start := time.Now()
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Printf("\rTimer: %v", time.Since(start))
			case <-done:
				return
			}
		}

	}()

	for i := 0; i < counter; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Please Enter a Integer")
			continue
		}
		if input < 0 || input > 100 {
			fmt.Println("Please Enter a number between 1-100 ")
			continue
		}
		if randomNum == input {
			fmt.Println("\nCongratulations you have guessed the correct number in", i+1, "tries!!!")
			return
		} else {
			fmt.Println("\nYou have guessed incorrect number!!")
			if input < randomNum {
				fmt.Println("The number is greater than the number you have entered!!")
				fmt.Println("You have ", counter-i-1, " tries left")

			} else {
				fmt.Println("Your number is smaller than the number you have enetered!!")
				fmt.Println("You have ", counter-i-1, " tries left")

			}
		}

	}
	done <- true
	ticker.Stop()
	elasped := time.Since(start)
	fmt.Println("\nGAME OVER!\nTHE CORRECT ANSWER WAS:", randomNum)
	fmt.Println("Total Time it took you:", elasped)
}

func playagain() bool {
	for {
		fmt.Println("Do you want to play again (y/n)")
		var again string
		_, err := fmt.Scan(&again)
		if err != nil {
			fmt.Println("Please Enter Y or N")
			continue
		}
		if again == "y" || again == "Y" {
			return true
		} else {
			fmt.Println("Alright! See you again!!")
			return false
		}
	}
}
