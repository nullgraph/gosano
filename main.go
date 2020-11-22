package main

import (
	"encoding/base64"
	"fmt"
	crypt "gosano/crypto"
)

func main() {
	fmt.Println("=== gosano project ===")

	// set1.Problem6("set1/6.txt")

	// plaintext := "April is the cruelest month, breeding\nLilacs out of the dead land, mixing\nMemory and desire, stirring\n Dull roots with spring rain.\nWinter kept us warm, covering\nEarth in forgetful snow, feeding\nA little life with dried tubers."
	// key := "TSELIOT"

	ciphertextString := "FSM3JSVvPSdzMSQsbzcmJiAgLDwgdD4qIj0neHQxNyksKz06NE8AICM1NyBlIzw7dDs1ZTghKnQwNiQoaSM1OjdpbCQmLD09IkYEKjk7ITxsKCEwdDcgPyA9MXhzNjggPSY9PSJGaQshOD9lPiYgICdzMiU9J3QnIzclJyh0JjIsImdFAz09MSk7bz8xIzFsPDx0IzI3IWVvNzslID4gITNeFiQ+PSd0PT1lKiY9MzEnIzklbyc6PDJgaSkxMTcsIi5FFXQ/LDg9IzF0PywqLG8jPSctbC09PTE3ZTg8LTEmIGs="
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextString)
	if err != nil {
		panic("file wasn't base64 encoded")
	}
	fmt.Println(ciphertext, len(ciphertext))
	// fmt.Println(crypt.GuessRepeatedXORKeyLength(ciphertext))
	crypt.DecryptRepeatedXOR(ciphertext)
}
