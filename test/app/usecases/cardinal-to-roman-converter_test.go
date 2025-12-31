package test

import (
	"strings"
	"testing"

	usecases "github.com/matheusadami/roman-numerals-converter/internal/app/usecases"
)

func TestExecuteMustReturnErrorWhenCardinalNumberIsNegative(t *testing.T) {
	// Given
	useCase := usecases.NewCardinalToRomanConverterUseCase()

	// When
	cardinalNumber := -1
	_, err := useCase.Execute(cardinalNumber)

	// Then
	expectedErrorMessage := "execute: cardinal number cannot be negative"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestExecuteMustReturnErrorWhenCardinalNumberIsGreaterThanMaxAllowed(t *testing.T) {
	// Given
	useCase := usecases.NewCardinalToRomanConverterUseCase()

	// When
	cardinalNumber := 100001
	_, err := useCase.Execute(cardinalNumber)

	// Then
	expectedErrorMessage := "execute: cardinal number cannot be greater than 100000"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("expected error message '%s', but got '%s'", expectedErrorMessage, err.Error())
	}
}

func TestExecuteMustReturnDefaultRomanNumberPhraseWhenCardinalNumberIsZero(t *testing.T) {
	// Given
	useCase := usecases.NewCardinalToRomanConverterUseCase()

	// When
	cardinalNumber := 0
	romanNumberPhrase, err := useCase.Execute(cardinalNumber)

	// Then
	if err != nil {
		t.Fatalf("unexpected error: %s", err.Error())
	}

	expectedPhrase := "N"
	if romanNumberPhrase != expectedPhrase {
		t.Errorf("expected roman number phrase '%s', but got '%s'", expectedPhrase, romanNumberPhrase)
	}
}
