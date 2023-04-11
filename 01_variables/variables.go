package main

import "fmt"

func maps() {
	m := make(map[string]int)
	fmt.Println("Empty String-Int Map - {} / size : {}", m, len(m))

	m["k1"] = 1
	m["k2"] = 2
	m["k3"] = 3

	fmt.Println("Filled String-Int Map - {} / size : {}", m, len(m))

	delete(m, "k2")
	fmt.Println("Deleted K2 from Map - {}", m)

	value_from_map, has_value := m["k2"]
	fmt.Println("Is K2 Present:", has_value)
	fmt.Println("Value From Map:", value_from_map)

	// Create and Initialize
	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Initialized Map : {}", m2)
}

func slices() {
	s := make([]string, 3)
	s[0] = "Hello"
	s[1] = "World"
	s[2] = "Element1"

	fmt.Println("String Filled Slice - {}", s)
	s = append(s, "Element2", "Element3")
	fmt.Println("Appended Element2 to Slice - {}", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copied Slice - {}", c)

	view_slice := s[2:5] // would be the same as s[2:]
	fmt.Println("View Slice - {}", view_slice)
}

func arrays() {
	var arr [3]string
	fmt.Println("String Empty Array - {}", arr)

	arr[0] = "Hello"
	arr[1] = "World"
	arr[2] = "Array"

	fmt.Println("String Filled Array - {}", arr)

	initialized_arr := [3]string{"This", "Array", "Initialized"}
	fmt.Println("String Initialized Array - {}", initialized_arr)
}

func usingShortHand() {
	str := "Hello World"
	fmt.Println("using_var_shorthand: {}", str)
}

func usingVar() {
	var str string
	str = "Hello World"
	fmt.Println("using_var method1 : {}", str)

	var str2 = "Hello World"
	fmt.Println("using_var method2 : {}", str2)

	var a, b int = 1, 2
	fmt.Println("using_var Int A : {} | Int B : {}", a, b)
}

func main() {

	fmt.Println("###### Using Var ######")
	fmt.Println()
	usingVar()

	fmt.Println("###### Using Assign Shorthand ######")
	fmt.Println()
	usingShortHand()

	fmt.Println("###### Arrays ######")
	fmt.Println()
	arrays()

	fmt.Println("###### Slices ######")
	fmt.Println()
	slices()

	fmt.Println("###### Maps ######")
	fmt.Println()
	maps()
}
