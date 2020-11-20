package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrequencyCount(t *testing.T) {
	s := "Cooking MC's like a pound of bacon."
	want := map[string]int{
		" ": 6, "'": 1, ",": 0, ".": 1, "a": 2, "b": 1, "c": 1, "d": 1, "e": 1, "f": 1, "g": 1, "h": 0, "i": 2, "j": 0, "k": 2, "l": 1, "m": 0, "n": 3, "o": 5, "p": 1, "q": 0, "r": 0, "s": 1, "t": 0, "u": 1, "v": 0, "w": 0, "x": 0, "y": 0, "z": 0,
	}
	got := FrequencyCount(s)
	assert.Equal(t, got, want)
}

func TestChi2Probability(t *testing.T) {
	s1 := "Cooking MC's like a pound of bacon."
	s2 := "Cooking MC's like a pound of bacon.$"
	assert.Less(t, Chi2Probability(s1), Chi2Probability(s2))
}
