package main

import (
	"log"
)

type Person struct {
	ID int
}

// tip use range with index only - for index := range list { list[index] }
func findMaxPerson(list []Person) *Person {
	// TODO: implement
	return nil
}

func testFindMaxPerson() {
	// arrange
	listPerson := []Person{
		{ID: 1},
		{ID: 20},
		{ID: 3},
		{ID: 1},
	}

	// act
	max := findMaxPerson(listPerson)

	// assert
	if max.ID != 20 {
		log.Fatalf("Expected Person ID number to be 20 but it was %d", max.ID)
	}
}

func main() {
	testFindMaxPerson()
}
