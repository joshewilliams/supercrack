package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/joshewilliams/supercrack/ciphers"
	"github.com/joshewilliams/supercrack/util"
)

func help() {
	fmt.Println("supercrack is your almost one stop shop for crypto cracking help")
	fmt.Println("Usage: supercrack <cipher|help|exit> OPTIONS")
	fmt.Printf("\nSupported Ciphers:\n\ncaesar\ntransposition (simple)\n\n")
	fmt.Printf("OPTIONS:\n\n")
	fmt.Printf("ciphertext=    Ciphertext string to crack\n")
	fmt.Printf("file=          Input filename\n")
	fmt.Printf("output=        Output filename\n\n")
}

func main() {
	fmt.Printf("SUPERCRACK\n\n")

mainLoop:
	for {
		cmd := prompt.Input("supercrack > ", util.MainMenuCompleter)
		p := util.Parameters{}

		values := strings.Fields(cmd)

		switch strings.ToLower(values[0]) {
		case "caesar":
			p.Cipher = "caesar"
		case "transposition":
			p.Cipher = "transposition"
		case "help":
			help()
			continue mainLoop
		// Add advanced case to handle cipher specific options
		// call advanced -> for loop, return breaks out of it, then you continue from the end of that chunk here until you type run
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid cipher selected")
		}

		for {
			cmd = prompt.Input("supercrack > ", util.SetupCompleter)
			if strings.Contains(cmd, "ciphertext=") {
				p.Ciphertext = []byte(cmd[11:])
			}

			if strings.Contains(cmd, "file=") {
				p.File = cmd[5:]
				var err error
				p.Ciphertext, err = ioutil.ReadFile(p.File)
				if err != nil {
					fmt.Println("Bad filename")
					continue
				}
			}

			if strings.Contains(cmd, "output=") {
				p.Output = cmd[7:]
			}

			if strings.Contains(cmd, "run") {
				break
			}
		}

		fmt.Printf("Cracking %s cipher\n", p.Cipher)
		switch p.Cipher {
		case "caesar":
			ciphertext := string(p.Ciphertext)
			ciphers.Caesar(ciphertext).WriteOutput(p.Output)
		case "transposition":
			ciphertext := string(p.Ciphertext)
			ciphers.Transposition(ciphertext).WriteOutput(p.Output)
		default:
			fmt.Println("Bad input")
		}
	}
}
