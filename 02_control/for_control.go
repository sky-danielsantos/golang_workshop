package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	ID int
}

// Range Pitfall

func findMax(list []Person) *Person {
	init_person_min := Person{ID: -10000}
	var max *Person = &init_person_min
	for _, p := range list {
		if p.ID > max.ID {
			max = &p
		}
	}
	return max
}

func rangePitallPtr2() {
	max_person := findMax([]Person{
		{ID: 5},
		{ID: 2},
		{ID: 3},
		{ID: 4},
	})
	fmt.Println("Max Person ID - {}", max_person.ID)
}

// Range Copies Values into the same variable
func rangeIsACopy() {
	m := make(map[string]*int)
	for i := 0; i < 3; i++ {
		value := i
		m["key-"+strconv.Itoa(i)] = &value
	}
	for k, v := range m {
		fmt.Printf("Memory Addr from map value: %p vs from range: %p\n", m[k], &v)
	}
}

// Map Iteration
func iterateMap() {
	m := make(map[string]int)
	number_of_keys := 15
	for i := 0; i < number_of_keys; i++ {
		m["key-"+strconv.Itoa(i)] = i
	}

	// Find every even value and Multiply it's value by 100
	for k, v := range m {
		if v%2 == 0 {
			m[k] = v * 100
		}
	}
	fmt.Println("Map after multiplication - {}", m)

	/*
		for k := range m {
			if m[k] % 2 == 0 {
				m[k] = v * 100
			}
		}*/
}

//

// Slice iteration
func findNumber(value_arr []*int, number int) *int {
	for i := 0; i < len(value_arr); i++ {
		if value_arr[i] != nil && *value_arr[i] == number {
			return value_arr[i]
		}
	}
	return nil
}

func iterateSlice() {
	s := make([]*int, 16)
	slice_size := 10
	i := 2
	for i < 2+slice_size {
		int_value := i * 10
		s[i] = &int_value
		i++
	}

	// We will see the memory address pointers instead of values
	fmt.Println("Slice After Fill - {}", s)
	// Has a pointer which value is 20 ?
	fmt.Println("Has number 20 - {}", findNumber(s, 20) != nil)
}

//

// Simple Example
func forExample() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even -> {}", i)
		}
	}
}

func main() {

	//fmt.Println("###### Even Nummber Loops ######")
	//fmt.Println()
	//forExample()

	//fmt.Println("###### Iterate For ######")
	//fmt.Println()
	//iterateSlice()

	//fmt.Println("###### Iterate For ######")
	//fmt.Println()
	//iterateMap()

	//fmt.Println("###### Range Is A Copy ######")
	//fmt.Println()
	//rangeIsACopy()

	//fmt.Println("###### Range Pitfall 2 ######")
	//fmt.Println()
	//rangePitallPtr2()

}
