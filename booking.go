package main

import (
	"fmt"
	"strings"
)

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
	var userTickets uint // uint --> unsigned integer --> only positive numbers

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

	// Updating the remaining tickets after booking also type conversion in present here
	remainingTickets = remainingTickets - int(userTickets)
	fmt.Printf("The remaining number of tickets are: %d\n", remainingTickets)

	// ------------ ARRAYS AND SLICES ------------
	// Array --> fixed size , Slice --> dynamic size
	// Array declaration :: var <variable_name> = [size]<data_type> OR var a = [5] int{1,2,3,4,5}

	var bookings = [50]string{"abc", "xyz"}
	// OR var bookings [50]string --> empty array of size 50 we can define array in any of the ways

	var firstName string = "Archisman"
	var lastName string = "Das"
	//Adding elements to the array
	bookings[0] = firstName + " " + lastName
	bookings[1] = "John Doe"

	fmt.Printf("The whole bookings are: %v\n", bookings)
	fmt.Printf("The first booking is: %v\n", bookings[0])
	fmt.Printf("Array Type %T\n", bookings)
	fmt.Printf("Array Length %v\n", len(bookings))

	// Slices --> dynamic sized array
	var userBookings []string // empty slice
	userBookings = append(userBookings, firstName+" "+lastName)
	userBookings = append(userBookings, "John Doe")

	fmt.Printf("The whole user bookings are: %v\n", userBookings)
	fmt.Printf("The first user booking is: %v\n", userBookings[0])
	fmt.Printf("Slice Type %T\n", userBookings)
	fmt.Printf("Slice Length %v\n", len(userBookings))

	// ------------- LOOPS -------------
	// we only have for-loop in GO

	//Infinite Loop
	// for {
	// 	fmt.Println("Hello World")
	// }

	// Loop with condition --> similar to while loop
	for remainingTickets > 0 {
		fmt.Println("Remaining tickets are:", remainingTickets)
		remainingTickets--
		if remainingTickets == 0 {
			fmt.Println("All tickets are booked")
			break
		}
	}

	// Loop with counter --> classic for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("Counter:", i)
	}

	// Loop with range
	for i, booking := range bookings {
		fmt.Printf("Booking %d: %v\n", i+1, booking)
	}

	// for each loop
	var firstNames = []string{}
	for _, name := range firstNames { // _ is used to ignore the index value i.e we only want the value
		var usersnames = strings.Fields(name) // splits the string into a slice of strings
		var firstName = usersnames[0]

		firstNames = append(firstNames, firstName)
	}

	fmt.Println("The first names are:", firstNames)

	// Conditional Statements -- If-else , else-if statements in GO
	var age int
	fmt.Println("Enter your age:")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else if age < 0 {
		fmt.Println("Invalid age")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	// Switch Case in GO
	switch age {
	case 16:
		fmt.Println("You are 16 years old")
	case 18:
		fmt.Println("You are 18 years old")
	case 21:
		fmt.Println("You are 21 years old")
	default:
		fmt.Println("Your age is neither 16, 18 nor 21")
	}

	// User Input Validation
	var email string

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")

	if isValidName && isValidEmail {
		fmt.Println("User input is valid. Proceeding with booking...")
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
	}

	// Functions in GO
	var guest = "Alice"
	greetUser(guest)

}

func greetUser(name string) {
	fmt.Println("Hello", name)
}
