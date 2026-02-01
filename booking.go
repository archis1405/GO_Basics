package main

import "fmt"

// FMT --> format package, used to format input and output
// FMT is used for Print Messages,CollectUser Input , Write to a File , etc.

// NOTE --> The main function is the entry point of a Go program. and we can have only one main function in a package.
// main function - entry point of the program
func main() {
	fmt.Println("Welcome to the Booking Application ...")
	fmt.Println("Get your tickets in the Booking Application ...")

	// Creating a variable in GO
	var name string = "Archisman Das"
	var conferenceName string = "Go Conference"

	fmt.Println("Hello", name, "Welcome to the", conferenceName)

	/*
		GO Compile errors for better code quality
		1. Unused Variables --> all the variables created must be used in the code
		2. Unused Imports --> all the imported packages must be used in the code
	*/

	// constant values defined ... constants are similar like variables only difference is that their values cannot be changed
	const tickets int = 50

	//tickets = 70 // this will throw an error because we cannot change the value of a constant

	var remainingTickets int = 50
	//remainingTickets = remainingTickets - 1 // this is valid because remainingTickets is a variable and can change

	fmt.Println("The total number of tickets are:", tickets)
	fmt.Println("The remaining number of tickets are:", remainingTickets)

	//Formatting output using printf
	fmt.Printf("The total number of tickets are: %d\n", tickets)
	fmt.Printf("The remaining number of tickets are: %d\n", remainingTickets)

	// User Input and Scanning
	var userName string
	var userTickets int

	// Asking user for their name and number of tickets similar to java Scanner class
	fmt.Println("Enter your name:")
	fmt.Scan(&userName) // & --> address operator --> gives the memory address of the variable

	// Pointer --> values are stored in a memory , pointer stores the address of the variable i.e. location in the memory where the variable is stored
	fmt.Println(&remainingTickets) // prints the memory address of the variable remainingTickets
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %s for booking %d tickets. You will receive a confirmation email soon.\n", userName, userTickets)

	// Types of variables in GO
	/*
		1. string --> sequence of characters
		2. int --> whole numbers
		3. float --> decimal numbers
		4. bool --> true or false
	*/

	fmt.Printf("User Name is of type: %T\n", userName)
	fmt.Printf("User Tickets is of type: %T\n", userTickets)
}
