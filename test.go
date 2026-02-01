package main // all our code should be part of a package

import "fmt" // import fmt package for formatted I/O

// test function - entry point of the program
func test() {
	//fmt.Print  --> to use the Print function from fmt package
	fmt.Println("Hello, World!") // print Hello, World! to the console
}

// Which Package when ? --> https://pkg.go.dev/std

// go run <file-name>  --> to run a go file from the terminal (without creating an executable)
