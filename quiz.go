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

	// Checking the user is eligible to play the game or not
	fmt.Printf("What is your age: ")
	var age uint
	fmt.Scan(&age)

	// If elseif else statement
	if age >= 10 {
		fmt.Println("Yay! you can play the game.")
	} else {
		fmt.Println("Sorry, you cannot play the game.")
		// return -> stops the execution of the program and return a value
		// Here it will stop the execution of the main function and return a value
		return
	}

	// The following statements will only be executed if the user is eligible to play the game

	// Quiz questions

	// Question 01

	fmt.Printf("What is better, the RTX 3080 or the RTX 3090: ")
	var answer string
	var answer2 string

	// Scan -> only scans the first word entered by the user (before separated by a space)
	// We need to use multiple variables to get multiple words from the user
	// There are other ways to take multiple words from the user too. (Check them out in the future)
	fmt.Scan(&answer, &answer2)

	// String concatenation -> "RTX" + "3090" -> "RTX3090"
	// We would need to add a space between the	answer and answer2
	if answer+" "+answer2 == "RTX 3090" {
		fmt.Println("Correct!")
	} else if answer+" "+answer2 == "rtx 3090" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

	// Question 02

	fmt.Printf("How many cores does the Ryzen 9 3900x have: ")

	// We don't need to declare the variable again if we are using the same variable name
	// But here we want a variable with type int, so we need to declare another variable
	// (because we cannot change the data type of a variable once it is declared)
	var answer3 uint
	fmt.Scan(&answer3)

	if answer3 == 12 {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect!")
	}

}
