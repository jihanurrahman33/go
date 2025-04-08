package main

import "fmt"

func main() {
	var username string = "Nishak"
	fmt.Println(username)
	//  %T represent type of the variable
	fmt.Printf("Variable is of type: %T ",username)



	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	//  %T represent type of the variable
	fmt.Printf("Variable is of type: %T ",isLoggedIn)

	var smallValue int =10
	//  %V represent the placeholder of the variable value
	fmt.Printf("Value is: %v",smallValue)
}