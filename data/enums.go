package data

func GetLetterPosition(data rune) int {
	letters := "abcdefghiljklmnopqrstuvwxyz"
	pos := make(map[rune]int)
	for i, j := range letters {
		pos[j] = i
	}

	return pos[data]
}

func NewLetterPosition(data int) rune {
	letters := "abcdefghijklmnopqrstuvwxyz"
	pos := make(map[int]rune)
	for i, j := range letters {
		pos[i] = j
	}

	return pos[data]
}
