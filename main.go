package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Mode struct {
	Choice string
	Number string
}

func main() {
	mode := Mode{}
	maxAttempts := 0
	Text()
	choiceNumber := InputChoice(mode.Choice)
	retryTimes := SetDifficulty(choiceNumber)

	GuessNumber(choiceNumber, mode.Number, maxAttempts, retryTimes)
	var again string
	TryAgain(again)
}

func Text() {
	fmt.Println("\nWelcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.")
	time.Sleep(1 * time.Second)
	fmt.Println("\nPlease select the difficulty level:")
	fmt.Println("1. 😏 Easy (10 attempts)")
	fmt.Println("2. 🙂 Medium (5 attempts)")
	fmt.Println("3. 😤 Hard (3 attempts)")
	time.Sleep(1 * time.Second)
}

func InputChoice(choice string) string {
	fmt.Print("\nYour choice: ")
	fmt.Scanf("%s\n", &choice)
	return choice
}

func Hint(number string, rand int) string {
	if number == "hint" {
		fmt.Printf("The number is near %v\n", rand+7)
		time.Sleep(1 * time.Second)
	}
	if number == "exact" {
		fmt.Printf("The number is between %v and %v\n", rand-7, rand+7)
		time.Sleep(1 * time.Second)
	}
	return number
}

func GuessNumber(choice string, number string, attempt int, retry int) int {
	randomNumber := rand.Intn(100 + 1)
	for attempt < retry {
		fmt.Print("\nYour guess: ")
		fmt.Scanf("%s\n", &number)
		number = Hint(number, randomNumber)
		guessNumber, err := strconv.Atoi(number)
		if err != nil && number != "hint" && number != "exact" {
			fmt.Println("Please input numbers only.")
			return attempt
		}
		if guessNumber > 0 && guessNumber <= 100 {
			if randomNumber > guessNumber {
				fmt.Printf("Incorrect! The number is greater than %v.", guessNumber)
			} else if randomNumber < guessNumber {
				fmt.Printf("Incorrect! The number is less than %v.", guessNumber)
			} else {
				fmt.Printf("Congratulations! You've guessed the correct number with %v attempts.", attempt+1)
				score := FinalScore(choice, attempt, retry)
				fmt.Printf("\nYour final score is: %v\n", score)
				break
			}
			fmt.Println("\nRetries left:", retry-attempt-1)
		} else {
			attempt--
			fmt.Println("\nPlease enter number in range between 1 and 100!")
			time.Sleep(1 * time.Second)
		}
		attempt++
		time.Sleep(1 * time.Second)
	}
	if attempt == retry {
		fmt.Printf("\nAlready used attempts, the correct number is %v\n", randomNumber)
		time.Sleep(1 * time.Second)
	}
	return attempt
}

func FinalScore(choice string, attempt int, retry int) int {
	switch choice {
	case "1":
		return (retry - attempt) * 10
	case "2":
		return (retry - attempt) * 20
	case "3":
		return (retry - attempt) * 30
	default:
		return 0
	}
}

func SetDifficulty(choice string) int {
	switch choice {
	case "1":
		fmt.Println("Easy")
		fmt.Println("10 attempts left!")
		return 10
	case "2":
		fmt.Println("Medium")
		fmt.Println("5 attempts left!")
		return 5
	case "3":
		fmt.Println("Hard")
		fmt.Println("⚠️	Guess number carefully! 3 attempts left!")
		return 3
	default:
		fmt.Println("❌	Invalid choice, please try again!")
		time.Sleep(1 * time.Second)
		return -1
	}
}

func TryAgain(playAgain string) {
	time.Sleep(1 * time.Second)
	fmt.Print("\nWould you like to play again? (y/n) ")
	fmt.Scanf("%s\n", &playAgain)
	if playAgain == "y" || playAgain == "Y" {
		main()
	} else {
		fmt.Println("\nThank you for playing! Goodbye!")
		time.Sleep(1 * time.Second)
	}
}
