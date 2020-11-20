package main

import (
	"fmt"
	crypt "gosano/crypto"
	// set1 "gosano/set1"
)

func main() {
	fmt.Println("=== gosano project ===")
	// s := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	// fmt.Println(crypt.DecryptSingleXOR(s))
	// t := "Cooking MC's like a pound of bacon."
	// fmt.Println(crypt.FrequencyCount(t))
	// fmt.Println(crypt.Chi2Probability(t))
	// fmt.Println(set1.Problem5())
	s1 := "this is a test"
	s2 := "wokka wokka!!!"
	fmt.Println(s1)
	fmt.Println(s2)
	distance, _ := crypt.HammingDistance(s1, s2)
	fmt.Println(distance)

}
