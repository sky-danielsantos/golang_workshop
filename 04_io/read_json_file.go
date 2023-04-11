package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Address struct {
	Avenue      string `json:"avenue"`
	SuiteNumber int    `json:"suiteNumber"`
}

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Address *Address `json:"address"`
}

func (person *Person) ToString() string {
	return fmt.Sprintf("[Name: %s, Age: %d, Address: %v]", person.Name, person.Age, person.Address)
}

type Users struct {
	Users []Person `json:"users"`
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		log.Fatal("error opening json file: ", err)
	}
	defer jsonFile.Close()

	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("error reading json file: ", err)
	}
	var users Users
	if err := json.Unmarshal(jsonBytes, &users); err != nil {
		log.Fatal("error deserializing json: ", err)
	}

	for index := range users.Users {
		person := &users.Users[index]
		fmt.Println(person.ToString())
	}
}
