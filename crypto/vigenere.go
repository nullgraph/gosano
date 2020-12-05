package crypto

import "fmt"

var vigenereTableau = GenerateVignereTableau()
var firstLetter = 65 // ascii code of the first letter

// GenerateVignereTableau produces the Vigenere table.
// It's a 26x26 table, the entries are uppercase letters.
func GenerateVignereTableau() [][]byte {
	table := make([][]byte, 26) // 26 letters in the alphabet
	for i := 0; i < 26; i++ {
		table[i] = make([]byte, 26)
	}
	// make first row
	for i := 0; i < 26; i++ {
		table[0][i] = byte(i + firstLetter)
	}
	for j := 1; j < 26; j++ {
		for i := 0; i < 26; i++ {
			table[j][i] = table[j-1][(i+1)%26]
		}
	}
	return table
}

// PrintVigenereTableau pretty prints the Vigenere tableau
func PrintVigenereTableau() {
	for i, row := range vigenereTableau {
		for j := range row {
			fmt.Printf("%q ", vigenereTableau[i][j])
		}
		fmt.Println()
	}
}

// EncryptWithVigenere encrypts the plaintext with the key using the Vigenere tableau
func EncryptWithVigenere(plaintext, key []byte) []byte {
	var ciphertext []byte
	for i := range plaintext {
		row := int(plaintext[i]) - firstLetter
		col := int(key[i%len(key)]) - firstLetter
		ciphertext = append(ciphertext, vigenereTableau[row][col])
	}
	return ciphertext
}

// func GuessVignereKeySize(ciphertext []byte)
