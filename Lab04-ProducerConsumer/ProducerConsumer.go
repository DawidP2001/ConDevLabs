//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Dawid Pionk
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import "sync"

type Event struct {
	ID      int
	Message string
}

func producer(theChannel chan Event, numLoops int, sharedLock *sync.Mutex, semaphore chan bool, wg *sync.WaitGroup) {
	for i := 0; i < numLoops; i++ {
		event := Event{ID: i, Message: "Hello"}
		sharedLock.Lock()
		theChannel <- event
		semaphore <- true
		sharedLock.Unlock()
	}
	wg.Done()
}

func consumer(theChannel chan Event, numLoops int, sharedLock *sync.Mutex, semaphore chan bool, wg *sync.WaitGroup) {
	for i := 0; i < numLoops; i++ {
		<-semaphore
		sharedLock.Lock()
		event := <-theChannel
		sharedLock.Unlock()
		print(event.ID, " ", event.Message)
	}
	wg.Done()
}
func a(wg *sync.WaitGroup) {
	print("HEllo")
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	//numOfThreads := 100
	numLoops := 5
	theChannel := make(chan Event)
	semaphore := make(chan bool)
	var sharedLock sync.Mutex

	go producer(theChannel, numLoops, &sharedLock, semaphore, &wg)
	go consumer(theChannel, numLoops, &sharedLock, semaphore, &wg)
	wg.Wait()
}
