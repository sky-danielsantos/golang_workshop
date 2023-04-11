package main

import (
	"fmt"
	"strconv"
)

func checkIfKeyExists(m map[string]*int, key string) {
	if value, ok := m[key]; ok {
		fmt.Printf("Key %s exists with value %d\n", key, *value)
	} else {
		fmt.Printf("Key %s does not exists\n", key)
	}
}

func ifMapExample() {
	m := make(map[string]*int)
	for i := 0; i < 3; i++ {
		value := i
		m["key-"+strconv.Itoa(i)] = &value
	}

	checkIfKeyExists(m, "key-1")
	checkIfKeyExists(m, "key-foobar")
}

func main() {
	ifMapExample()
}
