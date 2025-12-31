package entities

import (
	"fmt"

	consts "github.com/matheusadami/roman-numerals-converter/internal/shared/consts"
	mappers "github.com/matheusadami/roman-numerals-converter/internal/shared/mappers"
)

type RomanNumberPhrase struct {
	cardinalNumber int
	phrase         string
}

func NewRomanNumberPhrase(cardinalNumber int) (*RomanNumberPhrase, error) {
	if cardinalNumber < 0 {
		return nil, fmt.Errorf("cardinal number cannot be negative")
	}

	if cardinalNumber > consts.MaxAllowedCardinalNumber {
		return nil, fmt.Errorf("cardinal number cannot be greater than %d", consts.MaxAllowedCardinalNumber)
	}

	return &RomanNumberPhrase{cardinalNumber: cardinalNumber}, nil
}

func (e *RomanNumberPhrase) GetPhrase() string {
	if e.cardinalNumber <= 0 {
		return consts.DefaultRomanNumberPhrase
	}

	cardinalNum := e.cardinalNumber
	outputPhrase := ""

	for _, entry := range mappers.CardinalIntoRomanNumeralsMapper {
		for cardinalNum >= entry.Value {
			outputPhrase += entry.Symbol
			cardinalNum -= entry.Value
		}
	}

	return outputPhrase
}
