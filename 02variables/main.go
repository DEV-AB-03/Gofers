package main

import "fmt"

var jwtToken = 300000

// LoginToken public since first character is capital
const LoginToken = "ghabbb"

func main() {
	var username string = "aniket"
	fmt.Println(username)
	fmt.Printf("Variable is  of type %T\n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is  of type %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is  of type %T\n", smallVal)

	// default values and some aliases
	var variable int
	fmt.Println(variable)
	fmt.Printf("Variable is  of type %T\n", variable)

	// implicit
	var website = "learncodeonline"
	fmt.Println(website)

	// no var style
	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)
}
