package crypto

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// Guess represents a guess for a decryption
type Guess struct {
	Key         string
	Probability float32
	Plaintext   string
}

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

// DecryptSingleXOR guesses the single character that was used to XOR the plaintext. It uses chi2 probability test.
func DecryptSingleXOR(ciphertext []byte) Guess {
	var guesses []Guess

	for i := 48; i <= 122; i++ {
		b2 := RepeatedBytes(byte(i), len(ciphertext))
		xored := string(FixedXOR(ciphertext, b2))
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
func RepeatedXOR(plaintext, key []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b ^ key[i%len(key)]
	}
	return ciphertext
}

// HammingDistance calculates the number of differing **bits** between two strings. We do it by XOR'ing the two strings and then count the number of 1s in the result.
func HammingDistance(b1, b2 []byte) (int, error) {
	if len(b1) != len(b2) {
		return 0, fmt.Errorf("Hamming distance of byte arrays of different lengths %d and %d", len(b1), len(b2))
	}
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

// keysizeGuess represents a guess for the keysize in a repeated XOR cipher; distance is the normalized Hamming distance of the first two chunks.
type keysizeGuess struct {
	keysize  int
	distance float32
}

// guessRepeatedXORKeyLength tries to guess the key size for a repeated XOR encryption
func guessRepeatedXORKeyLength(ciphertext []byte) []keysizeGuess {
	var distances []keysizeGuess
	for i := 2; i <= 40; i++ {
		chunk1 := ciphertext[0:i]
		chunk2 := ciphertext[i : 2*i]
		distance, err := HammingDistance(chunk1, chunk2)
		if err != nil {
			panic(err)
		}
		distances = append(distances, keysizeGuess{i, float32(distance) / float32(i)})
	}
	sort.Slice(distances,
		func(i, j int) bool { return distances[i].distance < distances[j].distance })
	return distances
}

// chunkCiphertextForRepeatedXOR vertically breaks ciphertext into chunks of bytes. There are `keysize` rows, each contains the bytes of i-th column. Each row's length is dependent on how long the ciphertext is; in fact, we don't even care that the matrix comes out to be irregular because each row will be treated as a separate ciphertext of SingleXOR to be broken later.
func chunkCiphertextForRepeatedXOR(ciphertext []byte, keysize int) [][]byte {
	chunks := make([][]byte, keysize)
	for i := range ciphertext {
		chunks[i%keysize] = append(chunks[i%keysize], ciphertext[i])
	}
	return chunks
}

// DecryptRepeatedXOR decrypts a plaintext that was encrypted with the repeated XOR method.
func DecryptRepeatedXOR(ciphertext []byte) Guess {
	// choosing numberOfKeysizeGuesses is an art, adjust this parameter as needed
	numberOfKeysizeGuesses := 3
	keysizes := guessRepeatedXORKeyLength(ciphertext)
	// fmt.Println(keysizes)

	var guesses []Guess
	for i := 0; i < numberOfKeysizeGuesses; i++ {
		// fmt.Printf("i=%d keysizeGuess %v\n", i, keysizes[i])
		key := ""
		chunks := chunkCiphertextForRepeatedXOR(ciphertext, keysizes[i].keysize)
		for _, chunk := range chunks {
			// fmt.Println(chunk)
			columnGuess := DecryptSingleXOR(chunk)
			key += columnGuess.Key
		}
		plaintext := RepeatedXOR(ciphertext, []byte(key))
		prob := Chi2Probability(string(plaintext))
		guess := Guess{Key: key, Plaintext: string(plaintext), Probability: prob}
		guesses = append(guesses, guess)
	}

	sort.Slice(guesses,
		func(i, j int) bool { return guesses[i].Probability < guesses[j].Probability })

	return guesses[0]
}
