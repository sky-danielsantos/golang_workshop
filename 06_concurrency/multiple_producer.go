package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(channel chan<- string, producerID int, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		channel <- fmt.Sprintf("producer %d : %d", producerID, i)
	}
	wg.Done()
}
func consumer(channel <-chan string, done chan bool) {
	for {
		msg, more := <-channel
		if more {
			fmt.Println("Consumer: " + msg)
		} else {
			fmt.Println("Consumer: channel is closed finishing out")
			done <- true
			return // We exit the function
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	strChannel := make(chan string, 3)
	done := make(chan bool)
	numberOfProducer := 3
	wg := sync.WaitGroup{}
	for i := 0; i < numberOfProducer; i++ {
		go producer(strChannel, i, &wg)
		wg.Add(1)
	}
	go consumer(strChannel, done)
	wg.Wait()
	// Producers have finished so we can close the channel
	close(strChannel)

	<-done
}
