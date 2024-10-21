//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 14/10/2024
// Modified by: Dawid Pionk
// Description:
// A solution to the producer consumer problem
//--------------------------------------------

package main

import "sync"

// An struct for the message for producers and consumers
type Event struct {
	ID      int
	Message string
}

func producer(theChannel chan Event, numLoops int, wg *sync.WaitGroup) {
	for i := 0; i < numLoops; i++ {
		event := Event{ID: i, Message: "Hello"}
		theChannel <- event // Event put in channel
	}
	wg.Done()
}

func consumer(theChannel chan Event, numLoops int, wg *sync.WaitGroup) {
	for i := 0; i < numLoops; i++ {
		event := <-theChannel // Event taken out of channel
		print(event.ID, " ", event.Message)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	numLoops := 5
	theChannel := make(chan Event)

	go producer(theChannel, numLoops, &wg)
	go consumer(theChannel, numLoops, &wg)
	wg.Wait()
}
