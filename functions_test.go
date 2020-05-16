package main

import (
	"testing"
)

var erroMessage string = "Valor esperado %v, mas o valor do resultado foi %v\n"

func TestRandomNumbers(t *testing.T) {

	expectedLength := 12

	value := genrateRandomNumbers()
	valueLength := len(value)

	if expectedLength != valueLength {
		t.Errorf(erroMessage, expectedLength, valueLength)
	}

	gernerateCnpj()
}
