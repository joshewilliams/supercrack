package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joshewilliams/supercrack/classical"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	fmt.Printf("SUPERCRACK\n\n")

	// Flags
	var cipher string
	var ciphertext string
	var file string
	var output string

	flag.StringVar(&cipher, "c", "", "Cipher to crack")
	flag.StringVar(&file, "f", "", "Ciphertext file")
	flag.StringVar(&ciphertext, "t", "", "Ciphertext string")
	flag.StringVar(&output, "o", "", "Output file")
	flag.Parse()

	switch cipher {
	case "caesar":
		fmt.Println("Caesar Cipher Cracker")
		classical.Caesar(ciphertext).WriteOutput(output)
	case "transposition":
		fmt.Println("Transposition Cipher Cracker")
		classical.Transposition(ciphertext).WriteOutput(output)
	default:
		fmt.Println("Bad cipher name")
	}
}
