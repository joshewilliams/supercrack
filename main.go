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

func main() {
	fmt.Printf("SUPERCRACK\n\n")

mainLoop:
	for {
		cmd := prompt.Input("supercrack> ", util.Completer)
		p := util.Parameters{}

		values := strings.Fields(cmd)

		switch strings.ToLower(values[0]) {
		case "caesar":
			p.Cipher = "caesar"
		case "transposition":
			p.Cipher = "transposition"
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Invalid cipher selected")
		}

		for _, v := range values[1:] {
			if strings.Contains(v, "ciphertext=") {
				p.Ciphertext = []byte(v[11:])
			}

			if strings.Contains(v, "file=") {
				p.File = v[5:]
				var err error
				p.Ciphertext, err = ioutil.ReadFile(p.File)
				if err != nil {
					fmt.Println("Bad filename")
					continue mainLoop
				}
			}

			if strings.Contains(v, "output=") {
				p.Output = v[7:]
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
		}
		// Flags
		// var cipher string
		// var ciphertext string
		// var file string
		// var output string

		// flag.StringVar(&cipher, "c", "", "Cipher to crack")
		// flag.StringVar(&file, "f", "", "Ciphertext file")
		// flag.StringVar(&ciphertext, "t", "", "Ciphertext string")
		// flag.StringVar(&output, "o", "", "Output file")
		// flag.Parse()

		// switch cipher {
		// case "caesar":
		// 	fmt.Println("Caesar Cipher Cracker")
		// 	ciphers.Caesar(ciphertext).WriteOutput(output)
		// case "transposition":
		// 	fmt.Println("Transposition Cipher Cracker")
		// 	ciphers.Transposition(ciphertext).WriteOutput(output)
		// default:
		// 	fmt.Println("Bad cipher name")
		// }
	}
}
