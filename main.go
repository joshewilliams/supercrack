package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/joshewilliams/supercrack/classical"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func writeOutput(outputFile string, r []string) {
	var resultsString string
	for _, i := range r {
		resultsString += fmt.Sprintf("%s\n", i)
	}
	if outputFile != "" {
		err := ioutil.WriteFile(outputFile, []byte(resultsString), 0644)
		check(err)
	} else {
		fmt.Println(resultsString)
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
		writeOutput(output, classical.Caesar(ciphertext))
	default:
		fmt.Println("Bad cipher name")
	}
}
