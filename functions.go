package main

import (
	"math/rand"
	"strconv"
)

func genrateRandomNumbers() []int {

	numbers := make([]int, 12)

	for i := 0; i < 8; i++ {
		randNumber := rand.Intn(9)
		numbers[i] = randNumber
	}

	numbers[11] = 1

	return numbers
}

func mapingVector(numbers []int) [2][]int {

	var size int = len(numbers)
	var weight [2][]int

	weight[0] = numbers
	weight[1] = make([]int, size)

	var firstWeigth []int
	var secondWeigth []int

	if size == 12 {
		firstWeigth = weight[0][:4]
		secondWeigth = weight[0][4:]
	} else {
		firstWeigth = weight[0][:5]
		secondWeigth = weight[0][5:]
	}

	for i := 0; i < len(firstWeigth); i++ {
		weight[1][i] = len(firstWeigth) + 1 - i
	}

	for i := 0; i < len(secondWeigth); i++ {
		index := len(firstWeigth) + i
		weight[1][index] = 9 - i
	}

	return weight
}

func multiply(values [2][]int) []int {
	result := make([]int, len(values[0]))
	length := len(values[0])
	for i := 0; i < length; i++ {
		result[i] = values[0][i] * values[1][i]
	}

	return result
}

func sum(values []int) int {
	total := 0

	for _, number := range values {
		total += number
	}

	return total
}

func division(value int) int {
	return value % 11
}

func generateDigit(mapVector [2][]int) []int {
	m1 := multiply(mapVector)
	resultSum := sum(m1[:])
	divisionResult := division(resultSum)
	newValues := make([]int, len(m1), 14)
	newValues = append(mapVector[0][:])

	if divisionResult < 2 {
		newValues = append(newValues, 0)
		return newValues
	}
	digit := 11 - divisionResult
	newValues = append(newValues, digit)
	return newValues
}

func returnCnpjString(values []int) string {
	cnpj := ""
	for _, value := range values {
		cnpj += strconv.Itoa(value)
	}
	return cnpj
}

func gernerateCnpj() string {

	randomNumber := genrateRandomNumbers()
	mapV := mapingVector(randomNumber)
	withFirstDigit := generateDigit(mapV)
	mapV2 := mapingVector(withFirstDigit)
	completeCnpj := generateDigit(mapV2)

	return returnCnpjString(completeCnpj)
}

func arrayCpnj() [3]string {

	var cnpjs [3]string
	for i := 0; i < 3; i++ {
		cnpjs[i] = gernerateCnpj()
	}
	return cnpjs
}
