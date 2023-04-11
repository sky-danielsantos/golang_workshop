package main

import "fmt"

type Addres struct {
	avenue      string
	suiteNumber *int
}

type Person struct {
	name    string
	age     int
	address Addres
}

func changeSuiteNumnber(p Person, number int) {
	*p.address.suiteNumber = number
}

func main() {
	suiteNumber := 10
	person1 := Person{name: "Person 1", age: 10, address: Addres{avenue: "Avenue 1", suiteNumber: &suiteNumber}}
	person2 := person1

	fmt.Printf("Person 1 %v | Person 1 Suite Number : %d\n", person1, *person1.address.suiteNumber)
	fmt.Printf("Person 2 %v | Person 2 Suite Number : %d\n", person2, *person2.address.suiteNumber)

	changeSuiteNumnber(person1, 20000)

	fmt.Printf("Person 1 %v | Person 1 Suite Number : %d\n", person1, *person1.address.suiteNumber)
	fmt.Printf("Person 2 %v | Person 2 Suite Number : %d\n", person2, *person2.address.suiteNumber)

	fmt.Printf("Person 1 Address Struct Memory %p\n", &person1.address)
	fmt.Printf("Person 2 Address Struct Memory %p\n", &person2.address)
}
