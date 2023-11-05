package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	toEncrypt, encoding, message := getInput()
	if encoding == "1" {
		fmt.Println(rot13Cryption(message, toEncrypt))
		// Command to call the encryption/decription function 1 and output the final code
	} else if encoding == "2" {
		fmt.Println(reverseString(message))
		// Command to call the encryption/decription function 2 and output the final code
	} else {
		fmt.Println(customEncDec(message, toEncrypt))
		// Command to call the encryption/decription function 3 and output the final code
	}

}

// INPUT AND OUTPUT
func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Println("Welcome to the Cypher Tool!")
	// Command for output of the welcome phrase
	fmt.Println("Select operation (1/2):\n 1. Encrypt. \n 2. Decrypt.")
	var SelectOp int
	fmt.Scan(&SelectOp)
	// Command for choose Encrypt or Decrypt
	if SelectOp == 1 {
		toEncrypt = true
	} else if SelectOp == 2 {
		toEncrypt = false
	} else {
		fmt.Println("Incorrect input, please try again")
		getInput()
	}
	// Logic by which the toEncrypt variable is assigned true or false, also if the user enters incorrect data, the program is started again
	fmt.Println("Select cypher (1/2/3):\n 1. ROT13. \n 2. Reverse \n 3. CustomCipher")
	fmt.Scan(&encoding)
	if encoding == "1" || encoding == "2" || encoding == "3" {
		fmt.Println("Enter the message:")
		message = getMessage()
		// function call for text input
	} else {
		fmt.Println("Incorrect input, please try again")
		getInput()
	}
	return toEncrypt, encoding, message
	// Returns the entered data to the main function
}

func getMessage() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return scanner.Text()
	}
	return ""
	// Used the bafio library and os to enter the code with spaces, since fmt.Scan does not allow for this
}

// REVERSE ALPHABET ENCRYPTION AND DECRYPTION
// Spliting input in to runes and reversing them with ReverseAlphabetValue
func reverseString(input string) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Reverse each rune using the ReverseAlphabetValue()
	for i := 0; i < len(runes); i++ {
		//skip uppercases and other characters beside a-z
		if 'a' <= runes[i] && runes[i] <= 'z' {
			runes[i] = ReverseAlphabetValue(runes[i])
		}
	}

	// Convert the reversed runes into to a string
	reversedString := string(runes)
	return reversedString
}

// Reversing runes
func ReverseAlphabetValue(ch rune) rune {
	reversing := rune('z' - (ch - 'a'))
	return reversing
}

// ROT13 ENCRYPTION / DECRYPTION
func rot13Cryption(input string, toEncrypt bool) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Reverse each rune using the shiftBy function
	for i := 0; i < len(runes); i++ {
		// Skip characters other than 'a' to 'z'
		if 'a' <= runes[i] && runes[i] <= 'z' {
			runes[i] = shiftBy(runes[i], toEncrypt)
		}
	}

	// Convert the reversed runes into a string
	reversedString := string(runes)
	return reversedString
}

func shiftBy(r rune, toEncrypt bool) rune {
	step := 13
	if !toEncrypt {
		step = -step // If decryption, reverse the step
	}
	index := int(r - 'a')
	// Shifting index
	shiftedIndex := (index + step + 26) % 26 // Ensure the result is non-negative
	// Shifting rune
	shiftedRune := rune('a' + shiftedIndex)
	return shiftedRune
}

// CUSTOM ENCRYPTION AND DECRYPTION
func customEncDec(input string, toEncrypt bool) string {
	if toEncrypt == true {
		// Reverse alphabet encryption, then ROT13 encryption
		encryptedString := reverseString(input)
		encryptedString = rot13Cryption(encryptedString, true)
		return encryptedString
	} else {
		//ROT13 decryption, then reverse alphabet decryption
		decryptedString := rot13Cryption(input, false)
		decryptedString = reverseString(decryptedString)
		return decryptedString
	}
}