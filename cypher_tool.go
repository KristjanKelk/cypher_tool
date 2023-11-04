package main

import "fmt"

func main() {
	fmt.Println(getInput())
}

func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Println("Welcome to the Cypher Tool!")
	fmt.Println("Select operation (1/2):\n 1. Encrypt. \n 2. Decrypt.")
	var SelectOp int
	fmt.Scan(&SelectOp)
	if SelectOp == 1 {
		toEncrypt = true
	} else if SelectOp == 2 {
		toEncrypt = true
	}
	fmt.Println("Select cypher (1/2):\n 1. ROT13. \n 2. Reverse")
	fmt.Scan(&encoding)
	fmt.Println("Enter the message:")
	fmt.Scan(&message)
	return toEncrypt, encoding, message
}
