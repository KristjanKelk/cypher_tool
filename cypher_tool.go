package main

import "fmt"

func main() {
	fmt.Println(getInput())
}

func getInput() (toEncrypt bool, encoding string, message string) {
	fmt.Println("Welcome to the Cypher Tool!")
	// Command for output of the welcome phrase
	fmt.Println("Select operation (1/2):\n 1. Encrypt. \n 2. Decrypt.")
	var SelectOp int
	fmt.Scan(&SelectOp)
	if SelectOp == 1 {
		toEncrypt = true
	} else if SelectOp == 2 {
		toEncrypt = false
	}
	fmt.Println("Select cypher (1/2):\n 1. ROT13. \n 2. Reverse")
	fmt.Scan(&encoding)
	fmt.Println("Enter the message:")
	fmt.Scan(&message)
	return toEncrypt, encoding, message
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
func rot13Encryption(input string) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Reverse each rune using the ReverseAlphabetValue()
	for i := 0; i < len(runes); i++ {
		//skip uppercases and other characters beside a-z
		if 'a' <= runes[i] && runes[i] <= 'z' {
			runes[i] = shiftBy(runes[i])
		}
	}

	// Convert the reversed runes into to a string
	reversedString := string(runes)
	return reversedString
}

func shiftBy(r rune) rune {
	step := 13
	index := int(r - 'a')
	// shifting index
	shiftedIndex := (index + step) % 26
	//shifting rune
	shiftedRune := rune('a' + shiftedIndex)
	return shiftedRune
}
