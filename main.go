package main

import (
	"flag"
	"fmt"

	"github.com/joshewilliams/supercrack/classical"
)

func main() {
	fmt.Printf("SUPERCRACK\n\n")

	// Flags
	var cipher string
	var ciphertext string
	var file string

	flag.StringVar(&cipher, "c", "", "Cipher to crack")
	flag.StringVar(&file, "f", "", "Ciphertext file")
	flag.StringVar(&ciphertext, "t", "", "Ciphertext string")
	flag.Parse()

	switch cipher {
	case "caesar":
		fmt.Println("Caesar Cipher Cracker")
		results := classical.Caesar(ciphertext)
		for _, i := range results {
			fmt.Println(i)
		}
	default:
		fmt.Println("Bad cipher name")
	}
}
