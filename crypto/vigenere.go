package crypto

import (
	"fmt"
	"sort"
)

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

// GuessVignereKeySize tries to guess the keysize by "Hamming distance"
// It's basically the number of different letters normalized by length,
// the smaller the distance is, the more likely the two byte arrays are the same
func GuessVignereKeySize(ciphertext []byte) []keysizeGuess {
	var guesses []keysizeGuess
	for i := 3; i < 13; i++ {
		distance := 0
		for j := 0; j < i; j++ {
			distance += int(ciphertext[j] ^ ciphertext[i+j])
		}
		// hamming, _ := HammingDistance(ciphertext[0:i], ciphertext[i:i*2])
		// fmt.Println(i, float32(hamming)/float32(i), float32(distance)/float32(i))
		guesses = append(guesses, keysizeGuess{i, float32(distance) / float32(i)})
	}
	sort.Slice(guesses,
		func(i, j int) bool { return guesses[i].distance < guesses[j].distance })
	// fmt.Println(guesses)
	return guesses
}

// BreakVigenere breaks the Vignere ciphertext given
func BreakVigenere(ciphertext []byte) {
	numKeySizeGuesses := 3
	keysizeGuesses := GuessVignereKeySize(ciphertext)
	for i := 0; i < numKeySizeGuesses; i++ {
		fmt.Println(keysizeGuesses[i])
	}
}

// Modulus is not Remainder, i.e., it will return the least positive
// representation of the equivalence class of `a` mod `b`, not the remainder of
// `a` divided by `b`. There's probably no huge loss of performance here, it's
// expected that the compiler will inline this function.
func Modulus(a, b int) int {
	return (a%b + b) % b
}

// Rot is Ceasar shift, but it can also be done with the Vigenere Tableau
func Rot(plaintext []byte, key int) []byte {
	var ciphertext []byte
	for i := range plaintext {
		letter := int(plaintext[i]) - firstLetter + key
		ciphertext = append(
			ciphertext,
			byte(Modulus(letter, 26)+firstLetter))
	}
	return ciphertext
}

// BreakCeasar attempts to break the Ceasar cipher by chi2 probability testing
func BreakCeasar(ciphertext []byte) []Guess {
	var guesses []Guess
	for i := 0; i < 26; i++ {
		english := string(Rot(ciphertext, i))
		prob := Chi2Probability(english)
		key := string(byte(firstLetter + i))
		guesses = append(guesses, Guess{key, prob, english})
	}
	sort.Slice(
		guesses,
		func(i, j int) bool { return guesses[i].Probability < guesses[j].Probability })
	return guesses
}
