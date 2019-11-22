package classical

import (
	"math"
	"strings"
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

func Transposition(input string) string {
	return transpositionDecrypt(input, 2)
}
