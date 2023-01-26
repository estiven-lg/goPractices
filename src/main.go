package main

import (
	"fmt"
	"strings"
)

func isPalindrome(word string) bool {
	lowerWord := string(strings.ToLower(word))
	var reverseWord string

	for i := int(len(word) - 1); i >= 0; i-- {
		reverseWord += string(lowerWord[i])
	}

	return reverseWord == lowerWord
}

func main() {
	word := "Aibofobia"
	if isPalindrome(word) {
		fmt.Printf("'%s' es palanindroma \n", word)
	} else {
		fmt.Printf("'%s' no es palanindroma \n", word)
	}
}
