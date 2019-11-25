package ciphers

import (
	"fmt"
	"unicode"

	"github.com/joshewilliams/supercrack/util"
)

// Caesar function attempts to bruteforce a string believed to be encrypted by the Caesar cipher
func Caesar(ciphertext string) util.ResultsStringSlice {
	var results []string
	for i := 0; i < 26; i++ {
		var tmp []rune
		for _, j := range ciphertext {
			if !unicode.In(j, unicode.Letter) {
				tmp = append(tmp, j)
			} else if unicode.IsUpper(j) {
				tmp = append(tmp, util.SetLetterByPositionUpper((util.GetLetterPositionUpper(j)+i)%26))
			} else {
				tmp = append(tmp, util.SetLetterByPositionLower((util.GetLetterPositionLower(j)+i)%26))
			}
		}
		results = append(results, fmt.Sprintf("%d: %s", i, string(tmp)))
	}

	return results
}
