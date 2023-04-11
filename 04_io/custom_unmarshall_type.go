package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

var jsonText = []byte(`
{
  "name": "foobar",
  "creation_time": "2020-11-30T14:20:28.000Z"
}`)

type Person struct {
	Name     string    `json:"name"`
	Creation time.Time `json:"creation_time"`
}

// Example With Custom Time

var jsonTextWithCustomDate = []byte(`
{
  "name": "foobar",
  "creation_time": "2020-11-30"
}`)

type CustomTime struct {
	time.Time
}
type PersonWithCustomTime struct {
	Name     string     `json:"name"`
	Creation CustomTime `json:"creation_time"`
}

const CUSTOM_LAYOUT = "2006-01-02"

func (ct *CustomTime) UnmarshalJSON(bytes []byte) error {
	s := strings.Trim(string(bytes), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return nil
	}
	var err error
	ct.Time, err = time.Parse(CUSTOM_LAYOUT, s)
	if err != nil {
		return fmt.Errorf("failed to parse time from json: %v", err)
	}
	return nil
}

func main() {
	var person Person
	if err := json.Unmarshal(jsonText, &person); err != nil {
		log.Fatal("error on reading json: ", err)
	}
	fmt.Printf("Person unmarshall - %v\n", person)

	if err := json.Unmarshal(jsonTextWithCustomDate, &person); err != nil {
		log.Printf("expected error on reading json: %v", err)
	}

	var persnWithCustomTime PersonWithCustomTime
	if err := json.Unmarshal(jsonTextWithCustomDate, &persnWithCustomTime); err != nil {
		log.Fatal("error on reading json: ", err)
	}
	fmt.Printf("Person With Custom Time unmarshall - %v\n", persnWithCustomTime)
}
