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
