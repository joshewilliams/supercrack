package classical

import (
	"fmt"
	"unicode"

	"github.com/joshewilliams/supercrack/data"
)

func Caesar(ciphertext string) []string {
	var results []string
	for i := 0; i < 26; i++ {
		var tmp []rune
		for _, j := range ciphertext {
			if !unicode.In(j, unicode.Letter) {
				tmp = append(tmp, j)
			} else if unicode.IsUpper(j) {
				tmp = append(tmp, data.SetLetterByPositionUpper((data.GetLetterPositionUpper(j)+i)%26))
			} else {
				tmp = append(tmp, data.SetLetterByPositionLower((data.GetLetterPositionLower(j)+i)%26))
			}
		}
		results = append(results, fmt.Sprintf("%d: %s", i, string(tmp)))
	}

	return results
}
