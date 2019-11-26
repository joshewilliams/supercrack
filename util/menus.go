package util

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

func help() {
	fmt.Println("supercrack is your almost one stop shop for crypto cracking help")
	fmt.Println("Usage: supercrack <cipher|help|exit> OPTIONS")
	fmt.Printf("\nSupported Ciphers:\n\ncaesar\ntransposition (simple)\n\n")
	fmt.Printf("GENERAL OPTIONS:\n\n")
	fmt.Printf("ciphertext=    Ciphertext string to crack\n")
	fmt.Printf("file=          Input filename\n")
	fmt.Printf("output=        Output filename\n\n")
}

func MainMenu(p *Parameters) {
menuLoop:
	for {
		cmd := prompt.Input("supercrack > ", MainMenuCompleter, prompt.OptionPrefixTextColor(prompt.Red))

		switch strings.ToLower(cmd) {
		case "caesar":
			p.Cipher = "caesar"
			break menuLoop
		case "transposition":
			p.Cipher = "transposition"
			break menuLoop
		case "help":
			help()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid cipher selected")
			MainMenu(p)
		}
	}
}

func GeneralOptionsMenu(p *Parameters) {
menuLoop:
	for {
		cmd := prompt.Input("supercrack > ", GeneralOptionsCompleter, prompt.OptionPrefixTextColor(prompt.Red))
		splitCmd := strings.SplitN(cmd, "=", 2)

		switch splitCmd[0] {
		case "ciphertext":
			p.Ciphertext = []byte(splitCmd[1])
		case "file":
			p.File = splitCmd[1]
			var err error
			p.Ciphertext, err = ioutil.ReadFile(p.File)
			if err != nil {
				fmt.Println("Bad filename")
				continue menuLoop
			}
		case "output":
			p.Output = splitCmd[1]
		case "info":
			fmt.Println(p)
		case "run":
			break menuLoop
		case "hex":
			var err error
			len, err := hex.Decode(p.Ciphertext, p.Ciphertext)
			if err != nil {
				fmt.Println("Input not hex encoded")
				continue menuLoop
			}
			p.Ciphertext = p.Ciphertext[:len]
		case "b64":
			var err error
			len, err := base64.StdEncoding.Decode(p.Ciphertext, p.Ciphertext)
			if err != nil {
				fmt.Println("Input not b64 encoded")
				continue menuLoop
			}
			p.Ciphertext = p.Ciphertext[:len]
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Bad input")
			continue menuLoop
		}

	}
}
