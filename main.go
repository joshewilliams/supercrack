package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/joshewilliams/supercrack/classical"
)

type resultsStringSlice []string
type resultsString string

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func (r resultsStringSlice) writeOutput(outputFile string) {
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

func (r resultsString) writeOutput(outputFile string) {
	if outputFile != "" {
		err := ioutil.WriteFile(outputFile, []byte(r), 0644)
		check(err)
	} else {
		fmt.Println(r)
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
		resultsStringSlice(classical.Caesar(ciphertext)).writeOutput(output)
	case "transposition":
		fmt.Println("Transposition Cipher Cracker")
		resultsString(classical.Transposition(ciphertext)).writeOutput(output)
	default:
		fmt.Println("Bad cipher name")
	}
}
