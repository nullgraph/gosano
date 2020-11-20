package crypto

import (
	"strings"
)

// This file will only deal with the English language and will ignore all other complications like Unicode, runes, normalization, etc as warned in https://blog.golang.org/strings

// the last 3 data points, for , . and ' are from https://en.wikipedia.org/wiki/English_punctuation
var englishFreqs = map[string]float32{
	"a": 0.08167,
	"b": 0.01492,
	"c": 0.02782,
	"d": 0.04253,
	"e": 0.12702,
	"f": 0.02228,
	"g": 0.02015,
	"h": 0.06094,
	"i": 0.06966,
	"j": 0.00153,
	"k": 0.00772,
	"l": 0.04025,
	"m": 0.02406,
	"n": 0.06749,
	"o": 0.07507,
	"p": 0.01929,
	"q": 0.00095,
	"r": 0.05987,
	"s": 0.06327,
	"t": 0.09056,
	"u": 0.02758,
	"v": 0.00978,
	"w": 0.02360,
	"x": 0.00150,
	"y": 0.01974,
	"z": 0.00074,
	" ": 0.19181,
	",": 0.06130,
	".": 0.06530,
	"'": 0.02430,
}

// FrequencyCount counts the number of times a character in the alphabet appears in a string
func FrequencyCount(s string) map[string]int {
	freqs := make(map[string]int)
	for k := range englishFreqs {
		freqs[k] = strings.Count(s, k)
	}
	return freqs
}

// Chi2Probability calculates the probability that the string is English using chi2 testing. The lower the probability is, the more likely the text is English.
// See: https://crypto.stackexchange.com/a/30259/11248
func Chi2Probability(s string) float32 {
	// lowercase the string because the alphabet doesn't care about capitalization
	s = strings.ToLower(s)

	frequencies := FrequencyCount(s)

	// length of string ignoring stuff not in the alphabet
	length := 0
	for _, n := range frequencies {
		length += n
	}

	chi2 := float32(0)
	for _, c := range s {
		// fmt.Printf("%v %s %T\n", c, string(c), c)
		char := string(c)
		if freq, ok := englishFreqs[char]; ok {
			expected := float32(length) * freq
			difference := float32(frequencies[char]) - expected
			chi2 += difference * difference / expected
		} else { // character not in alphabet, give it a very high probability
			chi2 += 100
		}
	}

	return chi2
}
