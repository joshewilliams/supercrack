package data

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

type ResultsStringSlice []string

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
