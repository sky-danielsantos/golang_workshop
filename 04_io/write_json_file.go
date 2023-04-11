package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Address struct {
	Avenue      string `json:"avenue"`
	SuiteNumber int    `json:"suiteNumber"`
}

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address,omitempty"`
}

func (person *Person) ToString() string {
	return fmt.Sprintf("[Name: %s, Age: %d, Address: %v]", person.Name, person.Age, person.Address)
}

type Users struct {
	Users []Person `json:"users"`
}

func main() {
	address1 := Address{Avenue: "Avenue 1", SuiteNumber: 4}
	persons := []Person{
		{Name: "Person 1", Age: 10, Address: &address1},
		{Name: "Person 2", Age: 10},
	}
	users := Users{Users: persons}
	contentJson, err := json.MarshalIndent(users, "", "")
	if err != nil {
		log.Fatal("failed on json marshall: ", err)
	}
	ioutil.WriteFile("users.json", contentJson, 0644)
}
