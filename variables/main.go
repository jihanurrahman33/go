package main

import "fmt"
const LoginToken string ="Nishak"//public variable
//making the first letter capital means im making it public

func main() {
	var username string = "Nishak"
	fmt.Println(username)
	//  %T represent type of the variable
	fmt.Printf("Variable is of type: %T\n ",username)



	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	//  %T represent type of the variable
	fmt.Printf("Variable is of type: %T\n ",isLoggedIn)

	var smallValue int =1000000000000000000
	//  %V represent the placeholder of the variable value
	fmt.Printf("Value is: %v\n",smallValue)

	var secondSmallValue uint8 =255
	//  %T represent the type of variable 
	fmt.Printf("Value is: %T\n",secondSmallValue)


	var smallFloatValue float32 =255.32
	//  %T represent the type of variable 
	fmt.Printf("Value is: %T\n",smallFloatValue)


	//default value
	var anotherVariable int
	fmt.Println(anotherVariable)

//implicit type
	var website="google.com"
	fmt.Println(website)
	
	//no var style
	numberOfUser:=30000
	fmt.Println(numberOfUser)


	fmt.Println(LoginToken)
	fmt.Printf("Value is: %T\n",LoginToken)
}