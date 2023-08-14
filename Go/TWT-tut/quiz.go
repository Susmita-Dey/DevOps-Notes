package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz game!")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
