package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/matheusadami/roman-numerals-converter/internal/domain/entities"
)

func main() {
	var cardinalNumber int
	flag.IntVar(&cardinalNumber, "number", 0, "Cardinal number to be converted to Roman numeral")
	flag.Parse()

	romanPhrase, err := entities.NewRomanNumberPhrase(cardinalNumber)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Roman Numeral:", romanPhrase.GetPhrase())
}
