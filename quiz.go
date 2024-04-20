// We need to have a package name on top of every go file (it is mandatory)
// This is the main package
// This is the entry point of the application
// We can have only one main package in a project
// We can have multiple packages in a project
// We can have multiple files in a package (files are also called as modules)
package main

// Importing the required packages
// fmt - format package
// This allows us to print the output to the console and to read the input from the console
// This gives access to some tools build into the standard Go library.
import "fmt"

// This is the main function
// It is the entry point of the application
// It is the first function to be executed when the application is running (This is a executable function)
func main() {
	// P is capital in Println
	// Println -> Prints the output to the console and moves the cursor to the next line
	// No need to use semicolon at the end of the statement
	fmt.Println("Hello World")
}
