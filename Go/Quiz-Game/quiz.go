package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Quiz Game!")

	// Ask the user for their name
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello, %v! Welcome to the game\n", name)

	// Ask the user for their age
	var age uint
	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	// Check if the user is old enough to play the game
	if age >= 13 {
		fmt.Println("Yay! You are old enough to play this game!")
	} else {
		fmt.Println("Uh oh, you are too young to play this game!")
		return
	}

	// Rules of the game
	fmt.Println("\nHere are the rules:")
	fmt.Println("1. You will be asked 10 questions")
	fmt.Println("2. Each question will either have 4 options or no options")
	fmt.Println("3. You will have to enter the option number to answer the question if options are available")
	fmt.Println("4. Each question will have 10 points and there will be a penalty of 5 points for each wrong answer")

	score := 0 // score is initialized to 0

	// Start the game
	fmt.Println("\n\nLet's start the game!")
	fmt.Println("\nQuestion 1: What is the capital of India?")
	fmt.Println("1. New Delhi")
	fmt.Println("2. Mumbai")
	fmt.Println("3. Kolkata")
	fmt.Println("4. Chennai")

	// Declare the correct answer and take the user's answer as input
	var answer1 uint = 1
	var userAnswer1 uint = 0
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer1)

	// Check if the user's answer is correct
	if userAnswer1 == answer1 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("Option %v would be the correct answer\n", answer1)
	}

	fmt.Println("\nQuestion 2: If we combine Red and Hat, which company name it will become?")

	// Declare the correct answer and take the user's answer as input
	var answer2 string = "Red Hat"
	var userAnswer2a string
	var userAnswer2b string
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer2a, &userAnswer2b)

	// Check if the user's answer is correct
	if userAnswer2a+" "+userAnswer2b == answer2 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer2)
	}

	fmt.Println("\nQuestion 3: What is the largest mammal on Earth?")
	fmt.Println("1. African Elephant")
	fmt.Println("2. Blue Whale")
	fmt.Println("3. Giraffe")
	fmt.Println("4. Polar Bear")

	// Declare the correct answer and take the user's answer as input
	var answer3 uint = 2
	var userAnswer3 uint = 0
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer3)

	// Check if the user's answer is correct
	if userAnswer3 == answer3 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("Option %v would be the correct answer\n", answer3)
	}

	fmt.Println("\nQuestion 4: What is the chemical symbol for gold?")

	// Declare the correct answer and take the user's answer as input
	var answer4 string = "Au"
	var userAnswer4 string
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer4)

	// Check if the user's answer is correct
	if userAnswer4 == answer4 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer4)
	}

	fmt.Println("\nQuestion 5: Which planet is known as the 'Red Planet'?")
	fmt.Println("1. Venus")
	fmt.Println("2. Mars")
	fmt.Println("3. Jupiter")
	fmt.Println("4. Saturn")

	// Declare the correct answer and take the user's answer as input
	var answer5 uint = 2
	var userAnswer5 uint = 0
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer5)

	// Check if the user's answer is correct
	if userAnswer5 == answer5 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("Option %v would be the correct answer\n", answer5)
	}

	fmt.Println("\nQuestion 6: What is the tallest mountain in the world?")

	// Declare the correct answer and take the user's answer as input
	var answer6 string = "Mount Everest"
	var userAnswer6a, userAnswer6b string
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer6a, &userAnswer6b)

	// Check if the user's answer is correct
	if userAnswer6a+" "+userAnswer6b == answer6 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer6)
	}

	fmt.Println("\nQuestion 7: Who wrote the play 'Romeo and Juliet'?")
	fmt.Println("1. Charles Dickens")
	fmt.Println("2. Jane Austen")
	fmt.Println("3. William Shakespeare")
	fmt.Println("4. Mark Twain")

	// Declare the correct answer and take the user's answer as input
	var answer7 uint = 3
	var userAnswer7 uint = 0
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer7)

	// Check if the user's answer is correct
	if userAnswer7 == answer7 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer7)
	}

	fmt.Println("\nQuestion 8: Which famous artist painted the Mona Lisa?")

	// Declare the correct answer and take the user's answer as input
	var answer8 string = "Leonardo da Vinci"
	var userAnswer8a, userAnswer8b, userAnswer8c string
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer8a, &userAnswer8b, &userAnswer8c)

	// Check if the user's answer is correct
	if userAnswer8a+" "+userAnswer8b+" "+userAnswer8c == answer8 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer8)
	}

	fmt.Println("\nQuestion 9: What is the capital of France?")
	fmt.Println("1. Berlin")
	fmt.Println("2. London")
	fmt.Println("3. Madrid")
	fmt.Println("4. Paris")

	// Declare the correct answer and take the user's answer as input
	var answer9 uint = 4
	var userAnswer9 uint = 0
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer5)

	// Check if the user's answer is correct
	if userAnswer9 == answer9 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer9)
	}

	fmt.Println("\nQuestion 10: What is the largest organ in the human body?")

	// Declare the correct answer and take the user's answer as input
	var answer10 string = "Skin"
	var userAnswer10 string
	fmt.Printf("\nEnter your answer: ")
	fmt.Scan(&userAnswer10)

	// Check if the user's answer is correct
	if userAnswer10 == answer10 {
		score += 10
		fmt.Println("Correct answer!")
	} else {
		score -= 5
		fmt.Println("Wrong answer!")
		fmt.Printf("%v would be the correct answer\n", answer10)
	}

	fmt.Printf("\nYour score is %v\n", score)

	// Check if the user won or lost
	if score >= 80 {
		fmt.Println("Congratulations! You won!")
	} else {
		fmt.Println("Sorry, you lost!")
	}

	fmt.Println("Thanks for playing!")
}
