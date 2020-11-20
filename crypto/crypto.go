package crypto

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// FixedXOR takes two equal-length buffers and produces their XOR combination
func FixedXOR(b1 []byte, b2 []byte) []byte {
	// check equal length
	if len(b1) != len(b2) {
		panic(fmt.Sprintf("Error: XOR slices of length %d != %d", len(b1), len(b2)))
	}
	out := make([]byte, len(b1))
	for i := range b1 {
		out[i] = b1[i] ^ b2[i]
	}
	return out
}

// Guess represents a guess for a decryption
type Guess struct {
	Key         string
	Probability float32
	Plaintext   string
}

// DecryptSingleXOR guesses the single character that was used to XOR the plaintext. It uses chi2 probability test.
// cipher is a hex encoded string.
func DecryptSingleXOR(cipher string) Guess {
	var guesses []Guess

	b1, _ := hex.DecodeString(cipher)

	for i := 48; i <= 122; i++ {
		b2 := RepeatedBytes(byte(i), len(b1))
		xored := string(FixedXOR(b1, b2))
		prob := Chi2Probability(strings.ToLower(xored))
		guesses = append(guesses, Guess{string(i), prob, xored})
	}
	sort.Slice(guesses,
		func(i, j int) bool { return guesses[i].Probability < guesses[j].Probability })
	// fmt.Println(guesses)
	return guesses[0]
}

// EncryptSingleXOR encrypts the plaintext with a repetition of the key.
// The ciphertext is hex encoded
func EncryptSingleXOR(plaintext string, key byte) string {
	b1 := []byte(plaintext)
	b2 := RepeatedBytes(key, len(b1))
	encoded := hex.EncodeToString(FixedXOR(b1, b2))
	return encoded
}

// FindKeyForSingleXOR guesses the key that XOR plaintext into ciphertext
func FindKeyForSingleXOR(plaintext, ciphertext string) int {
	b1 := []byte(plaintext)
	for i := 48; i <= 122; i++ {
		b2 := RepeatedBytes(byte(i), len(b1))
		encoded := hex.EncodeToString(FixedXOR(b1, b2))
		if strings.Contains(ciphertext, encoded) {
			return i
		}
	}
	return -1
}

// RepeatedBytes returns a byte slice that consists of repeated byte `b` for `length` times.
func RepeatedBytes(b byte, length int) []byte {
	bytes := make([]byte, length)
	for j := range bytes {
		bytes[j] = b
	}
	return bytes
}

// RepeatedXOR xors the key repeatedly against the plaintext. The ciphertext is hex encoded.
func RepeatedXOR(plaintext, key string) string {
	blaintext := []byte(plaintext)
	bkey := []byte(key)

	bcipher := make([]byte, len(blaintext))
	for i, blain := range blaintext {
		bcipher[i] = blain ^ bkey[i%len(bkey)]
	}
	return hex.EncodeToString(bcipher)
}

// HammingDistance calculates the number of differing **bits** between two strings. We do it by XOR'ing the two strings and then count the number of 1s in the result.
func HammingDistance(s1, s2 string) (int, error) {
	if len(s1) != len(s2) {
		return 0, fmt.Errorf("Hamming distance of strings of different lengths %d and %d", len(s1), len(s2))
	}
	b1 := []byte(s1)
	b2 := []byte(s2)
	distance := 0
	for i := range b1 {
		xored := b1[i] ^ b2[i]
		for j := 0; j < 8; j++ {
			distance += int(xored & 1)
			xored = xored >> 1
		}
	}
	return distance, nil
}
