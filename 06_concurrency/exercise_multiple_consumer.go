package main

import (
	"fmt"
	"strconv"
	"time"
)

func producer(channel chan string, done chan bool) {
	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("ping " + strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
	close(channel)
	<-done
}
func consumer(channel chan string, done chan bool) {
	for {
		msg, more := <-channel
		if more {
			fmt.Println("Consumer: " + msg)
		} else {
			fmt.Println("Consumer: channel is closed finishing out")
			done <- true
			return // We exit the function
		}
	}
}

func main() {
	strChannel := make(chan string)
	done := make(chan bool)
	go consumer(strChannel, done)
	producer(strChannel, done)

}
