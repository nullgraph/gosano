package crypto

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptWithVigenere(t *testing.T) {
	plaintext := "TO BE OR NOT TO BE THAT IS THE QUESTION"
	plaintext = strings.Join(strings.Fields(plaintext), "")
	// fmt.Println(plaintext)
	key := "RELATIONS"
	want := "KSMEHZBBLKSMEMPOGAJXSEJCSFLZSY"
	cipheretext := EncryptWithVigenere([]byte(plaintext), []byte(key))
	got := string(cipheretext)
	// fmt.Println(got)
	assert.Equal(t, want, got)
}

func TestModulus(t *testing.T) {
	assert.Equal(t, 3, Modulus(3, 26))
	assert.Equal(t, 23, Modulus(-3, 26))
	assert.Equal(t, 0, Modulus(26, 26))
	assert.Equal(t, 1, Modulus(27, 26))
}

func TestRot(t *testing.T) {
	plaintext := "THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG"
	plaintext = strings.Join(strings.Fields(plaintext), "")
	want := "WKHTXLFNEURZQIRAMXPSVRYHUWKHODCBGRJ"
	got := string(Rot([]byte(plaintext), 3))
	assert.Equal(t, want, got)

	got = string(Rot([]byte(plaintext), 26))
	assert.Equal(t, plaintext, got)

	want = "GURDHVPXOEBJASBKWHZCFBIREGURYNMLQBT"
	got = string(Rot([]byte(plaintext), 13))
	assert.Equal(t, want, got)

	want = "QEBNRFZHYOLTKCLUGRJMPLSBOQEBIXWVALD"
	got = string(Rot([]byte(plaintext), -3))
	assert.Equal(t, want, got)
}
