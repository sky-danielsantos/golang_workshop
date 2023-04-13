package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("please provide a csv file to convert")
	}
	fileName := os.Args[1]
	fmt.Println(ReadAndConvert2Json(fileName, ","))
}
