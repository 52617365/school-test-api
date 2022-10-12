package exampleData

import (
	"math/rand"
)

func GenerateRandomFirstName() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfRandomFirstNames))
	randomFirstName := listOfRandomFirstNames[randomIndex]
	return randomFirstName
}
func GenerateRandomLastName() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfRandomLastNames))
	randomLastName := listOfRandomLastNames[randomIndex]
	return randomLastName
}

func GenerateRandomSocialSecurityNumberName() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfSocialSecurityNumbers))
	randomSocialSecurityNumber := listOfSocialSecurityNumbers[randomIndex]
	return randomSocialSecurityNumber
}
func GenerateRandomBloodPressure() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfBloodPressure))
	randomBloodPressure := listOfBloodPressure[randomIndex]
	return randomBloodPressure
}
func GenerateRandomWeight() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfWeights))
	randomWeight := listOfWeights[randomIndex]
	return randomWeight
}

func GenerateRandomHeight() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfHeights))
	randomHeight := listOfHeights[randomIndex]
	return randomHeight
}

func GenerateRandomAge() string {
	randomIndex := generateRandomNumberInbetweenNumbers(0, len(listOfAges))
	randomAges := listOfAges[randomIndex]
	return randomAges
}

func generateRandomNumberInbetweenNumbers(min int, max int) int {
	randomIndex := rand.Intn(max)
	return randomIndex
}
