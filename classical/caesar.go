package classical

import (
	"fmt"

	"github.com/joshewilliams/supercrack/data"
)

func Caesar(ciphertext string) []string {
	var results []string
	for i := 0; i < 26; i++ {
		var tmp []rune
		for _, j := range ciphertext {
			tmp = append(tmp, data.NewLetterPosition((data.GetLetterPosition(j)+i)%26))
		}
		results = append(results, fmt.Sprintf("%d: %s", i, string(tmp)))
	}

	return results
}
