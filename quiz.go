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
	fmt.Println("Welcome to my quiz game!")

	// Datatype cannot be changed for a variable once it is declared
	// var name string = "Tim"
	// You can also declare a variable without specifying the data type too. Go will automatically detect the data type
	// var name = "Tim"
	// name := "Tim"
	// Note -> In go, if you declare a variable, you have to use it. Otherwise, it will throw an error

	// Declaring a variable for name without assigning a value
	var name string

	// To get input from the user, we can use Scanln
	// Scan -> reads the input from the console and stores it in the variable and moves the cursor to the next line
	// Scanln -> reads the input from the console and stores it in the variable and does not move the cursor to the next line
	// & -> address of the variable (pointer) -> memory address of the variable in RAM
	// If the user enters a value in a different data type, it will give the default value of the data type of the variable
	fmt.Printf("What is your name: ") // We use Printf here to avoid going to the next line
	fmt.Scan(&name)

	// We can use printf to format the output (Printf -> Print formatted)
	// %v -> value
	// %d -> decimal
	// %f -> float
	// %s -> string
	// \n -> new line
	fmt.Printf("Hello %v, Welcome to the game!\n", name) // fmt.Println("Hello", name)

	// Some data type in GO
	// string -> "Hello"
	// int -> 10 (signed integer)
	// uint -> 10 (unsigned integer)
	// float64, float32 -> 10.5
	// bool -> true , false

}
