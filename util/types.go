package util

import (
	"fmt"
	"io/ioutil"
	"log"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ResultsStringSlice type used to simplify information output via STDOUT or specified file
type ResultsStringSlice []string

// WriteOutput function for the ResultsStringSlice type to simplify information output
func (r ResultsStringSlice) WriteOutput(outputFile string) {
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

// Parameters type to store/parse parameters used
type Parameters struct {
	Cipher     string
	Ciphertext []byte
	File       string
	Output     string
}
