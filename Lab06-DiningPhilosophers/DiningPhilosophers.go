//--------------------------------------------
// Author: Dawid Pionk
// Created on 21/10/24
// Description:
// A solution to the dinining philosophers problem
// Issues:
// None I hope
//--------------------------------------------

package main

import (
	"math/rand/v2"
	"sync"
	"time"
)

func think(myID int, thinkTime int) {
	seconds := rand.IntN(thinkTime) + 1
	print(myID, " is thinking!\n")
	time.Sleep(time.Duration(seconds) * time.Second)
}

func get_forks(philID int, forks map[int]chan bool, count int, semChannel chan bool) {
	semChannel <- true              // Semaphore wait
	forks[philID] <- true           // Fork left wait
	forks[(philID+1)%count] <- true // Fork right wait
}

func put_forks(philID int, forks map[int]chan bool, count int, semChannel chan bool) {
	<-forks[philID]           // fork left go
	<-forks[(philID+1)%count] // fork right go
	<-semChannel              // Sempaphore go
}
func eat(myId int, eatTime int) {
	seconds := rand.IntN(eatTime) + 1
	print(myId, " is chomping!\n")
	time.Sleep(time.Duration(seconds) * time.Second)
}

func philosopher(id int, wg *sync.WaitGroup, forks map[int]chan bool, thinkTime int, count int, eatTime int, semChannel chan bool) {
	for {
		think(id, thinkTime)
		get_forks(id, forks, count, semChannel)
		eat(id, eatTime)
		put_forks(id, forks, count, semChannel)
	}
	wg.Done() // Will never be reached since philosophers eat forever
}

func main() {
	totalRoutines := 5
	var wg sync.WaitGroup
	wg.Add(totalRoutines)

	count := 5     // Number of philosophers
	thinkTime := 3 // Max amount in seconds for how long they think
	eatTime := 5   // Max amount in seconds for how long they eat

	semChannel := make(chan bool, 4) // Makes cannel for semaphores
	forks := make(map[int]chan bool) // Makes a channel of channels
	for i := 0; i < 5; i++ {
		forks[i] = make(chan bool, 1) // Places channels inside forks
	}
	for i := range totalRoutines {
		// id comes from the i above
		go philosopher(i, &wg, forks, thinkTime, count, eatTime, semChannel)
	}
	wg.Wait()
}
