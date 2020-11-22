package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptSingleXOR(t *testing.T) {
	plaintext := "Cooking MC's like a pound of bacon"
	key := 88
	ciphertext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	got := EncryptSingleXOR(plaintext, byte(key))
	assert.Equal(t, ciphertext, got)
}

func TestFindKeyForSingleXOR(t *testing.T) {
	plaintext := "Cooking MC's like a pound of bacon"
	key := 88
	ciphertext := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	got := FindKeyForSingleXOR(plaintext, ciphertext)
	assert.Equal(t, key, got)
}

func TestRepeatedXOR(t *testing.T) {
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	ciphertext := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	got := RepeatedXOR(plaintext, key)
	assert.Equal(t, ciphertext, got)
}

func TestHammingDistance(t *testing.T) {
	// trying to calculate Hamming distance of strings of different lengths should produce error
	_, err := HammingDistance([]byte("abc"), []byte("abcde"))
	assert.Error(t, err)
	// success test
	s1 := "this is a test"
	s2 := "wokka wokka!!!"
	got, err := HammingDistance([]byte(s1), []byte(s2))
	want := 37
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestDecryptVigenere(t *testing.T) {
	// plaintext := "April is the cruelest month, breeding\nLilacs out of the dead land, mixing\nMemory and desire, stirring\n Dull roots with spring rain.\nWinter kept us warm, covering\nEarth in forgetful snow, feeding\nA little life with dried tubers."
	// key := "TSELIOT"
	// // encrypt with repeated XOR, decode hex and re-encode with base64
	// ciphertext := "FSM3JSVvPSdzMSQsbzcmJiAgLDwgdD4qIj0neHQxNyksKz06NE8AICM1NyBlIzw7dDs1ZTghKnQwNiQoaSM1OjdpbCQmLD09IkYEKjk7ITxsKCEwdDcgPyA9MXhzNjggPSY9PSJGaQshOD9lPiYgICdzMiU9J3QnIzclJyh0JjIsImdFAz09MSk7bz8xIzFsPDx0IzI3IWVvNzslID4gITNeFiQ+PSd0PT1lKiY9MzEnIzklbyc6PDJgaSkxMTcsIi5FFXQ/LDg9IzF0PywqLG8jPSctbC09PTE3ZTg8LTEmIGs="
}
