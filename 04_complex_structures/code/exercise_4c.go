package main

import "fmt"

var initialPets map[string]string = map[string]string{
	"fido":     "dog",
	"penelope": "horse",
	"nancy":    "cat",
}

func doesPetExist(petName string) bool {
	_, ok := initialPets[petName]

	return ok
}

func main() {
	fmt.Println(
		doesPetExist("fido"))
}
