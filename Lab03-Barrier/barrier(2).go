//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Dawid Pionk
// Description: A simple barrier using semaphors and mutexes
// Issues:
// None
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, theLock *sync.Mutex, semChan chan bool) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	theLock.Lock()
	*arrived++

	if *arrived == max {
		theLock.Unlock()
		semChan <- true
		<-semChan
	} else {
		theLock.Unlock()
		<-semChan
		semChan <- true
	}
	//we wait here until everyone has completed part A
	fmt.Println("Part B", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	semChan := make(chan bool)
	var theLock sync.Mutex
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, &theLock, semChan)
	}

	wg.Wait() //wait for everyone to finish before exiting
}
