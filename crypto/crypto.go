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

// DecryptSingleXOR guesses the single character that was used to XOR the plaintext.
// It uses chi2 probability test.
func DecryptSingleXOR(ciphertext []byte) []Guess {
	var guesses []Guess

	for i := 0; i <= 255; i++ {
		b2 := RepeatedBytes(byte(i), len(ciphertext))
		xored := string(FixedXOR(ciphertext, b2))
		prob := Chi2Probability(strings.ToLower(xored))
		guesses = append(guesses, Guess{string(i), prob, xored})
	}
	sort.Slice(guesses,
		func(i, j int) bool { return guesses[i].Probability < guesses[j].Probability })
	// fmt.Println(guesses)
	return guesses
}

// EncryptSingleXOR encrypts the plaintext with a repetition of the key.
// The ciphertext is hex encoded.
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

// HammingDistance calculates the number of differing **bits** between two byte arrays.
// We do it by one byte at a time by XOR'ing them and then count the number of 1s in the result.
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

// keysizeGuess represents a guess for the keysize in a repeated XOR cipher;
// distance is the normalized Hamming distance of the first two chunks.
type keysizeGuess struct {
	keysize  int
	distance float32
}

// guessRepeatedXORKeySize tries to guess the key size for a repeated XOR encryption
func guessRepeatedXORKeySize(ciphertext []byte) []keysizeGuess {
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

func alternativeGuessForRepeatedXORKeySize(ciphertext []byte) []keysizeGuess {
	var distances []keysizeGuess
	for i := 2; i <= 40; i++ {
		chunk1 := ciphertext[0:i]
		chunk2 := ciphertext[i : 2*i]
		chunk3 := ciphertext[2*i : 3*i]
		chunk4 := ciphertext[3*i : 4*i]
		distance1, err := HammingDistance(chunk1, chunk2)
		if err != nil {
			panic(err)
		}
		distance2, err := HammingDistance(chunk3, chunk4)
		if err != nil {
			panic(err)
		}
		distances = append(distances, keysizeGuess{i, float32(distance1) / float32(distance2)})
	}
	sort.Slice(distances,
		func(i, j int) bool { return distances[i].distance < distances[j].distance })
	return distances
}

// chunkCiphertextForRepeatedXOR vertically breaks ciphertext into chunks of bytes.
// There are `keysize` rows, each contains the bytes of i-th column.
// Each row's length is dependent on how long the ciphertext is;
// in fact, we don't even care that the matrix comes out to be irregular
// because each row will be treated as a separate ciphertext of SingleXOR
// to be broken later.
func chunkCiphertextForRepeatedXOR(ciphertext []byte, keysize int) [][]byte {
	chunks := make([][]byte, keysize)
	for i := range ciphertext {
		chunks[i%keysize] = append(chunks[i%keysize], ciphertext[i])
	}
	return chunks
}

// MakePermutationsFromBuckets produces permutations of the choices,
// i.e., each slot should contain all the choices for that slot.
// In total, there should be (numChoices)**(numSlots) elements in the result.
func MakePermutationsFromBuckets(elts [][]string, numSlots, slotIndex int) []string {
	if slotIndex == numSlots-1 {
		return elts[slotIndex]
	}
	laterElts := MakePermutationsFromBuckets(elts, numSlots, slotIndex+1)
	var outputs []string
	for _, choice := range elts[slotIndex] {
		for _, later := range laterElts {
			outputs = append(outputs, choice+later)
		}
	}
	return outputs
}

// DecryptRepeatedXOR decrypts a plaintext that was encrypted
// with the repeated XOR method.
// This is analogous to breaking Vigenere.
func DecryptRepeatedXOR(ciphertext []byte) Guess {
	// choosing numberOfKeysizeGuesses is an art,
	// adjust this parameter as needed
	numberOfKeysizeGuesses := 3
	// numberOfFixedXORGuesses is the allowed guesses for each FixedXOR
	// vertical chunks. Be careful to keep this number small or the total
	// number of guesses (numberOfFixedXORGuesses**numberOfKeysizeGuesses)
	// will blow up.
	// numberOfFixedXORGuesses := 3
	numberOfFixedXORGuesses := 2
	fmt.Println(numberOfKeysizeGuesses, numberOfFixedXORGuesses)

	// keysizes := guessRepeatedXORKeySize(ciphertext)
	// fmt.Println(keysizes)
	keysizes := alternativeGuessForRepeatedXORKeySize(ciphertext)
	fmt.Println(keysizes)

	var guesses []Guess
	for n := 0; n < numberOfKeysizeGuesses; n++ {
		fmt.Printf("n=%d keysizeGuess %v\n", n, keysizes[n])
		columnGuesses := make([][]Guess, keysizes[n].keysize)
		chunks := chunkCiphertextForRepeatedXOR(ciphertext, keysizes[n].keysize)
		for i, chunk := range chunks {
			// fmt.Println(chunk)
			columnGuesses[i] = DecryptSingleXOR(chunk)
			// fmt.Printf("   %v %v\n", columnGuess.Probability, columnGuess.Key)
			// key += columnGuess.Key
		}

		keyBuckets := make([][]string, keysizes[n].keysize)
		for i := range columnGuesses {
			for j := 0; j < numberOfFixedXORGuesses; j++ {
				keyBuckets[i] = append(keyBuckets[i], columnGuesses[i][j].Key)
			}
		}
		fmt.Println(keyBuckets)

		keys := MakePermutationsFromBuckets(keyBuckets, keysizes[n].keysize, 0)
		// fmt.Println(keys)
		fmt.Println("num keys", len(keys))

		for _, key := range keys {
			plaintext := RepeatedXOR(ciphertext, []byte(key))
			prob := Chi2Probability(string(plaintext))
			// fmt.Printf("key=%v prob=%v    ", key, prob)
			guess := Guess{Key: key, Plaintext: string(plaintext), Probability: prob}
			guesses = append(guesses, guess)
		}

	}

	sort.Slice(guesses,
		func(i, j int) bool { return guesses[i].Probability < guesses[j].Probability })

	return guesses[0]

	// return Guess{}
}
