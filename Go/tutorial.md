# Go Tutorial 
**(CC: TechWorld with Nana)**

```go
package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	//  Note that ":=" is only applied to vars and not const
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	// var bookings [50]string //array
	var bookings []string //slices

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n")

	// fmt.Println(conferenceName)
	// fmt.Println(&conferenceName)

	// var bookings = [50]string{"Nana","Nicole","Peter"}

	// bookings[0] = "Susmita"
	// bookings[1] = "Nana"

	// fmt.Println(bookings)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("\nEnter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		// Pointers in golang are also called as special variables

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice size: %v\n", len(bookings))

		fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	}

}
```

### For loop syntax
```go
for index, booking := range bookings {}
```

```go
for _, booking := range bookings {}
// here _ is known as a blank identifier to ignore a variable you don't want to use. So, with Go you need to make unused variables explicitly
```

## Functions in Go

```go
func func_name(param_name param_type) param_return_type {
	// some code
	return some_value
}
```

**Note: A Go function can return multiple values together.**

## Package Level Variables
- Defined at the top outside all functions
- They are accessed inside any of the functions
- And in all files, which are in the same package.

## Packages in Go

- Go programs are organized into packages
- A package is a collection of Go files

Command:
```bash
go run .
```
or,
```bash
go run main.go helper.go
```

Here main.go file is dependent on helper.go file. So, we need to run both of them. Alternatively, we can use the first command if you're inside the folder.

*Capitalizing fucntion name in the package will export that function*

## Maps in Go
- Maps unique keys to values
- You can retrieve the value by using its key later

## Structs
- Stands for "Structure"
- Can hold mixed data types
- Defining a structure(which fields) of the User Type 
- It is like a lightweight class, which e.g. doesn't support inheritance.

**Create a struct:**

Syntax -

```go
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
	isOptedInForNewsletter bool
}
```

Here the **type keyword** creates a new type, with the name you specify.

*"Create a type called "UserData" based on a struct of firstName, lastName,..."*

## Concurrency in Go

**Concurrency makes our program more efficient. In Golang, it's cheap and easy.**

To make your code concurrent, use the **go** keyword.

```go
go sendTicket()
```

## Waitgroup
- Waits for the launched goroutine to finish
- Package "sync" provides basic synchronization functionality
- **Add:** Sets the number of goroutines to wait for (increases the counter by the provided number)
- **Wait:** Blocks until the WaitGroup counter is 0
- **Done:** Decrements the WaitGroup counter by 1. So this is called by the goroutine to indicate that it's finished