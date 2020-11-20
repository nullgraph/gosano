package crypto

import (
	"fmt"
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
	_, err := HammingDistance("abc", "abcde")
	fmt.Println(err)
	assert.Error(t, err)
	// success test
	s1 := "this is a test"
	s2 := "wokka wokka!!!"
	fmt.Println(s1)
	fmt.Println(s2)
	got, err := HammingDistance(s1, s2)
	want := 37
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
