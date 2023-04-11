package main

import (
	"fmt"
	"log"
	"os"
)

// Simple Defer vs Defer Function
func readWithSimpleDefer() {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		log.Fatalf("error failed to open file %s with %s", "products.json", err)
	}
	defer jsonFile.Close() // What if Close fails ?
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		log.Fatal("failed to close file", err)
	}
}

func readWithBetterDefer() {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		log.Fatalf("error failed to open file %s with %s", "products.json", err)
	}
	defer CloseFile(jsonFile)
}

//

// Question: In Which order is this going to be executed ?
func deferOrder() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
}

//

// Defer Argument Evaluation is it lazy (delayed) or eager ?
var variableDeferExecution int

func deferExecution(argInt int) {
	fmt.Println("Defer variable execution :", argInt)
}
func incrementAndGetVarDefer() int {
	variableDeferExecution++
	return variableDeferExecution
}

// Question: When deferExecution does execute what is the int ?
func deferArgsEvalutation() {
	defer deferExecution(incrementAndGetVarDefer())
	fmt.Println("Defer Variable execution before on println before end of function ", incrementAndGetVarDefer())
}

//

func main() {
	deferOrder()
	deferArgsEvalutation()

	/*
		What if I have the following code

		for i := 0 ; i < 10; i++ {
			jsonFile, err := os.Open("products.json")
			if err != nil {
				log.Fatalf("error failed to open file %s with %s", "products.json", err)
			}
			defer jsonFile.Close() // What if Close fails ?
		}
	*/
}
