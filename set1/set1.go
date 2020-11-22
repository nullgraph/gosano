package set1

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	crypt "gosano/crypto"
	"io/ioutil"
	"os"
	"sort"
)

// Problem1 converts hex to base64.
// Link: https://cryptopals.com/sets/1/challenges/1
func Problem1() string {
	fmt.Println("Cryptopals Set1 Problem1")
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(s)
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	s64 := base64.StdEncoding.EncodeToString(b)
	return s64
}

// Problem2 takes two equal-length buffers and produces their XOR combination.
// Link: https://cryptopals.com/sets/1/challenges/2
func Problem2() string {
	b1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b2, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	out := crypt.FixedXOR(b1, b2)
	s := hex.EncodeToString(out)
	return s
}

// Problem3 decodes a string that was XOR'd against a single character.
// Link: https://cryptopals.com/sets/1/challenges/3
func Problem3() crypt.Guess {
	s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	guess := crypt.DecryptSingleXOR(s)
	return guess
}

// Problem4 detects a string which was single-XOR'd in a file.
func Problem4(filename string) crypt.Guess {
	var ciphers []string
	var guesses []crypt.Guess

	// read file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ciphers = append(ciphers, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, cipher := range ciphers {
		guess := crypt.DecryptSingleXOR(cipher)
		guesses = append(guesses, guess)
	}

	sort.Slice(guesses,
		func(i, j int) bool { return guesses[i].Probability < guesses[j].Probability })

	return guesses[0]
}

// ReverseProblem4 helped solve Problem 4. I was stuck with it for so long and so badly, and the main problem turned out that my key space was too small!!!
func ReverseProblem4() {
	var ciphers []string
	// read file
	file, err := os.Open("set1/4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ciphers = append(ciphers, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	plaintext := "Now that the party is jumping\n"

	for _, cipher := range ciphers {
		guess := crypt.FindKeyForSingleXOR(plaintext, cipher)
		if guess != -1 {
			fmt.Printf("key=%d, cipher=%s\n", guess, cipher)
		}
	}
}

// Problem5 implements repeating-key XOR
func Problem5() string {
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	key := "ICE"
	ciphertext := crypt.RepeatedXOR(plaintext, key)
	return ciphertext
}

// Problem6 breaks repeated XOR
// https://cryptopals.com/sets/1/challenges/6
func Problem6(filename string) string {
	// read file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("file %v not available", filename))
	}

	cipher, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		panic("file wasn't base64 encoded")
	}
	fmt.Println(cipher)

	return ""
}
