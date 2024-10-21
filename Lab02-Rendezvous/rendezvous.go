package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

//Global variables shared between functions --A BAD IDEA

func WorkWithRendezvous(wg *sync.WaitGroup, Num int, barrierSem chan bool, arrived *int) bool {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Part A", Num)
	//Rendezvous here
	*arrived++
	if *arrived == 5 {
		barrierSem <- true
		<-barrierSem
	} else {
		<-barrierSem
		barrierSem <- true
	}
	fmt.Println("PartB", Num)
	wg.Done()
	return true
}

func main() {
	var wg sync.WaitGroup
	barrierSem := make(chan bool)
	threadCount := 5
	arrived := 0
	wg.Add(threadCount)
	for N := range threadCount {
		go WorkWithRendezvous(&wg, N, barrierSem, &arrived)
	}
	wg.Wait() //wait here until everyone (10 go routines) is done

}
