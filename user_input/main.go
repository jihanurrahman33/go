package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome:="User input Practice"
	fmt.Println(welcome)

	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	//comma ok syntax,comma error syntax
	input ,_:=reader.ReadString('\n')
	fmt.Println("Your Name is: ",input)
	fmt.Printf("%v is %T",input,input)
}