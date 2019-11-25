package util

import "github.com/c-bata/go-prompt"

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		// Cipherlist
		{Text: "caesar", Description: "Caesar cipher"},
		{Text: "transposition", Description: "Transposition cipher (simple)"},
		{Text: "help", Description: "Help menu"},
		{Text: "exit", Description: "Exit supercrack"},
		//Parameters
		{Text: "ciphertext=", Description: "Ciphertext to crack"},
		{Text: "file=", Description: "Read input file"},
		{Text: "output=", Description: "Output file name"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
