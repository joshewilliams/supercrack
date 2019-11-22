package data

const (
	UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LOWER = "abcdefghijklmnopqrstuvwxyz"
)

// GetLetterPositionLower function determines a given rune's numeric position in the alphabet, lower case version
func GetLetterPositionLower(data rune) int {
	pos := make(map[rune]int)
	for i, j := range LOWER {
		pos[j] = i
	}

	return pos[data]
}

// SetLetterByPositionLower function returns the letter represented by a numeric position in the alphabet, lower case version
func SetLetterByPositionLower(data int) rune {
	pos := make(map[int]rune)
	for i, j := range LOWER {
		pos[i] = j
	}

	return pos[data]
}

// GetLetterPositionUpper determines a given rune's numeric position in the alphabet, upper case version
func GetLetterPositionUpper(data rune) int {
	pos := make(map[rune]int)
	for i, j := range UPPER {
		pos[j] = i
	}

	return pos[data]
}

// SetLetterByPositionLower function returns the letter represented by a numeric position in the alphabet, upper case version
func SetLetterByPositionUpper(data int) rune {
	pos := make(map[int]rune)
	for i, j := range UPPER {
		pos[i] = j
	}

	return pos[data]
}
