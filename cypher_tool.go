package main

import "fmt"

func main() {
	fmt.Println(getInput())
}

// INPUT AND OUTPUT
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
func rot13Cryption(input string, encryption bool) string {
	// Convert the string to a slice of runes
	runes := []rune(input)

	// Reverse each rune using the shiftBy function
	for i := 0; i < len(runes); i++ {
		// Skip characters other than 'a' to 'z'
		if 'a' <= runes[i] && runes[i] <= 'z' {
			runes[i] = shiftBy(runes[i], encryption)
		}
	}

	// Convert the reversed runes into a string
	reversedString := string(runes)
	return reversedString
}

func shiftBy(r rune, encryption bool) rune {
	step := 13
	if !encryption {
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
func customEncDec(input string, encryption bool) string {
	if encryption == true {
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
