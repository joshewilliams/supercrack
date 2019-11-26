package util

import "github.com/c-bata/go-prompt"

func MainMenuCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		// Cipherlist
		{Text: "caesar", Description: "Caesar cipher"},
		{Text: "transposition", Description: "Transposition cipher (simple)"},
		{Text: "xor", Description: "XOR"},
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
		{Text: "hexd", Description: "Ciphertext is hex encoded"},
		{Text: "hexe", Description: "Hex encode ciphertext"},
		{Text: "b64d", Description: "Ciphertext is b64 encoded"},
		{Text: "b64e", Description: "b64 encode ciphertext"},
		{Text: "exit", Description: "Exit this program"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
