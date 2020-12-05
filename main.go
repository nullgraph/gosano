package main

import (
	"fmt"
	"strings"

	crypt "gosano/crypto"
	// set1 "gosano/set1"
)

func main() {
	fmt.Println("=== gosano project ===")

	// set1.Problem6("set1/6.txt")

	// // encrypt with repeated XOR, decode hex and re-encode with base64
	// ciphertextString := "FSM3JSVvPSdzMSQsbzcmJiAgLDwgdD4qIj0neHQxNyksKz06NE8AICM1NyBlIzw7dDs1ZTghKnQwNiQoaSM1OjdpbCQmLD09IkYEKjk7ITxsKCEwdDcgPyA9MXhzNjggPSY9PSJGaQshOD9lPiYgICdzMiU9J3QnIzclJyh0JjIsImdFAz09MSk7bz8xIzFsPDx0IzI3IWVvNzslID4gITNeFiQ+PSd0PT1lKiY9MzEnIzklbyc6PDJgaSkxMTcsIi5FFXQ/LDg9IzF0PywqLG8jPSctbC09PTE3ZTg8LTEmIGs="
	// // now try to break it
	// ciphertext, _ := base64.StdEncoding.DecodeString(ciphertextString)
	// guess := crypt.DecryptRepeatedXOR(ciphertext)
	// fmt.Println(guess)

	plaintext := "TO BE OR NOT TO BE THAT IS THE QUESTION"
	plaintext = strings.Join(strings.Fields(plaintext), "")
	fmt.Println(plaintext)
	key := "RELATIONS"
	cipheretext := crypt.EncryptWithVigenere([]byte(plaintext), []byte(key))
	fmt.Println(string(cipheretext))
}
