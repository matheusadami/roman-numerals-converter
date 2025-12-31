package usecases

import (
	"fmt"

	"github.com/matheusadami/roman-numerals-converter/internal/domain/entities"
)

type CardinalToRomanConverterUseCase struct{}

func NewCardinalToRomanConverterUseCase() *CardinalToRomanConverterUseCase {
	return &CardinalToRomanConverterUseCase{}
}

func (uc *CardinalToRomanConverterUseCase) Execute(cardinalNumber int) (string, error) {
	romanNumberPhrase, err := entities.NewRomanNumberPhrase(cardinalNumber)
	if err != nil {
		return "", fmt.Errorf("execute: %v", err)
	}

	return romanNumberPhrase.GetPhrase(), nil
}
