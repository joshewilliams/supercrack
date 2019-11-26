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

		switch p.Cipher {
		case "caesar":
			ciphers.Caesar(string(p.Ciphertext)).WriteOutput(p.Output)
		case "transposition":
			ciphers.Transposition(string(p.Ciphertext)).WriteOutput(p.Output)
		case "xor":
			fmt.Printf("%x\n", ciphers.Xor(&p))
		default:
			fmt.Println("Bad input")
		}
	}
}
