package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var quit = false
	var count = 3
	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)
	fmt.Println("Number Gussing Game")
	for !quit {
		fmt.Println(count, " Life Left")
		fmt.Printf("Enter the number ")
		fmt.Scan(&userInput)
		if userInput == secretNumber {
			fmt.Println("You Won")
			quit = true
		} else if userInput > secretNumber {
			fmt.Println("Too High..")
			count--
		} else if userInput < secretNumber {
			fmt.Println("Too Low..")
			count--
		}
		if count == 0 {
			quit = true
			fmt.Println("You Loss")
		}
	}
}
