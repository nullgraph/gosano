package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	got := Problem1()
	assert.Equal(t, got, want)
}

func TestProblem2(t *testing.T) {
	want := "746865206b696420646f6e277420706c6179"
	got := Problem2()
	assert.Equal(t, got, want)
}

func TestProblem3(t *testing.T) {
	plaintext := "Cooking MC's like a pound of bacon"
	key := "X"
	guess := Problem3()
	assert.Equal(t, guess.Plaintext, plaintext)
	assert.Equal(t, guess.Key, key)
}

func TestProblem4(t *testing.T) {
	plaintext := "Now that the party is jumping\n"
	key := "5"
	guess := Problem4("4.txt")
	assert.Equal(t, guess.Plaintext, plaintext)
	assert.Equal(t, guess.Key, key)
}

func TestProblem5(t *testing.T) {
	ciphertext := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	got := Problem5()
	assert.Equal(t, got, ciphertext)
}
