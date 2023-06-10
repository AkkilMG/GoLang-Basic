package main

import (
	"fmt"
	"strings"
)

func caesarCipher(text string, shift int) string {
	return strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return 'A' + ((r - 'A' + int32(shift)) % 26)
		} else if r >= 'a' && r <= 'z' {
			return 'a' + ((r - 'a' + int32(shift)) % 26)
		}
		return r
	}, text)
}

func main() {
	var text string
	var shift int
	fmt.Scan(&text)
	fmt.Scan(&shift)
	encryptedText := caesarCipher(text, shift)

	fmt.Println("Original Text:", text)
	fmt.Println("Shift:", shift)
	fmt.Println("Encrypted Text:", encryptedText)
}
