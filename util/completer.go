package util

import "github.com/c-bata/go-prompt"

func MainMenuCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		// Cipherlist
		{Text: "caesar", Description: "Caesar cipher"},
		{Text: "transposition", Description: "Transposition cipher (simple)"},
		{Text: "help", Description: "Help menu"},
		{Text: "exit", Description: "Exit supercrack"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
func GeneralOptionsCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		//Parameters
		{Text: "ciphertext=", Description: "Ciphertext to crack"},
		{Text: "file=", Description: "Read input file"},
		{Text: "output=", Description: "Output file name"},
		{Text: "info", Description: "Print current parameters"},
		{Text: "run", Description: "Execute cracking"},
		{Text: "hex", Description: "Ciphertext is hex encoded"},
		{Text: "b64", Description: "Ciphertext is b64 encoded"},
		{Text: "exit", Description: "Exit this program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
