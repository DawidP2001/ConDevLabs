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
// Created on 7/10/2024
// Modified by: Dawid Pionk
// Description:
// A simple reusable barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, arrived *int, max int, wg *sync.WaitGroup, sharedLock *sync.Mutex, theChan chan bool, theChan2 chan bool) bool {
	for i := 0; i < 4; i++ {
		time.Sleep(time.Second)
		fmt.Println("Part A", goNum)
		//we wait here until everyone has completed part A
		sharedLock.Lock()
		*arrived++
		if *arrived == max { //last to arrive -signal others to go
			sharedLock.Unlock() //unlock before any potentially blocking code
			theChan <- true
			<-theChan
		} else { //not all here yet we wait until signal
			sharedLock.Unlock() //unlock before any potentially blocking code
			<-theChan
			theChan <- true //once we get through send signal to next routine to continue
		} //end of if-else
		// Wait here until everyone has left the previous barrier
		sharedLock.Lock()
		*arrived--
		if *arrived == 0 {
			sharedLock.Unlock()
			theChan2 <- true
			<-theChan2
		} else {
			sharedLock.Unlock()
			<-theChan2
			theChan2 <- true
		}
		fmt.Println("Part B", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 10
	arrived := 0
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	var theLock sync.Mutex
	theChan := make(chan bool) //use unbuffered channel in place of semaphore
	theChan2 := make(chan bool)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &arrived, totalRoutines, &wg, &theLock, theChan, theChan2)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
