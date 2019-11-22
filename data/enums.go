package data

const (
	UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER = "abcdefghijklmnopqrstuvwxyz"
)

func GetLetterPositionLower(data rune) int {
	pos := make(map[rune]int)
	for i, j := range LOWER {
		pos[j] = i
	}

	return pos[data]
}

func SetLetterPositionLower(data int) rune {
	pos := make(map[int]rune)
	for i, j := range LOWER {
		pos[i] = j
	}

	return pos[data]
}
func GetLetterPositionUpper(data rune) int {
	pos := make(map[rune]int)
	for i, j := range UPPER {
		pos[j] = i
	}

	return pos[data]
}

func SetLetterPositionUpper(data int) rune {
	pos := make(map[int]rune)
	for i, j := range UPPER {
		pos[i] = j
	}

	return pos[data]
}
