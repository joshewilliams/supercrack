package ciphers

import (
	"fmt"
	"math"
	"strings"

	"github.com/joshewilliams/supercrack/util"
)

func transpositionEncrypt(input string, key int) string {
	ciphertext := make([]string, key)
	for i := 0; i < len(input); i++ {
		ciphertext[i%key] = ciphertext[i%key] + string(input[i])
	}

	return strings.Join(ciphertext, "")
}

func transpositionDecrypt(input string, key int) string {
	key = int(math.Ceil(float64(len(input)) / float64(key)))
	cleartext := make([]string, int(key))
	for i := 0; i < len(input); i++ {
		cleartext[i%key] = cleartext[i%key] + string(input[i])
	}

	return strings.Join(cleartext, "")
}

// Transposition function bruteforces a given string believed to be encrypted via a simple transposition cipher
func Transposition(input string) util.ResultsStringSlice {
	var results util.ResultsStringSlice
	for i := 1; i < len(input); i++ {
		results = append(results, fmt.Sprintf("%d: %s", i, transpositionDecrypt(input, i)))
	}

	return results
}
