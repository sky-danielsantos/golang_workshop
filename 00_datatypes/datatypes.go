package main

import "fmt"

func zero_the_value(ival int) {
	ival = 0
}
func zero_the_value_ptr(ival *int) {
	*ival = 0
}

func pointers() {
	var some_int int = 10
	zero_the_value(some_int)
	fmt.Println("Int Value - {}", some_int)

	zero_the_value_ptr(&some_int)
	fmt.Println("Int Value - {}", some_int)

}

func manyInts() {
	var some_uint uint = 12
	fmt.Println("UInt - {}", some_uint)
	var some_uint_32 uint32 = 12
	fmt.Println("Uint32 - {}", some_uint_32)
	var some_int int = -12
	fmt.Println("Singed Int - {}", some_int)
}

func simpleTypes() {

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

}

func main() {

	fmt.Println("###### Simple Types ######")
	simpleTypes()

	fmt.Println("###### Many Ints ######")
	manyInts()

	fmt.Println("###### Pointer Ints ######")
	pointers()
}
