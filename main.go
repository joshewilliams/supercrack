package main

import (
	"fmt"

	"github.com/joshewilliams/supercrack/ciphers"
	"github.com/joshewilliams/supercrack/util"
)

func main() {
	for {
		p := util.Parameters{}

		util.MainMenu(&p)
		util.GeneralOptionsMenu(&p)

		fmt.Printf("Cracking %s cipher\n", p.Cipher)
		switch p.Cipher {
		case "caesar":
			ciphers.Caesar(string(p.Ciphertext)).WriteOutput(p.Output)
		case "transposition":
			ciphers.Transposition(string(p.Ciphertext)).WriteOutput(p.Output)
		default:
			fmt.Println("Bad input")
		}
	}
}
