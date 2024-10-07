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
// Description:
// A simple barrier implemented using mutex and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, theChan chan bool, theChan2 chan bool, counter *int32, max *int32, min *int32) bool {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Second)
		fmt.Println("Part A", goNum)
		atomic.AddInt32(counter, 1)
		if *counter == *max {
			theChan <- true
			<-theChan
		} else {
			<-theChan
			theChan <- true
		} //end of if-else
		atomic.AddInt32(counter, -1)
		if *counter == *min {
			theChan2 <- true
			<-theChan2
		} else {
			<-theChan2
			theChan2 <- true
		}
		fmt.Println("Part B", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	totalRoutines := 5
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	theChan := make(chan bool) //use unbuffered channel in place of semaphore
	theChan2 := make(chan bool)
	var counter int32
	var max int32 = 5
	var min int32 = 0
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, theChan, theChan2, &counter, &max, &min)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
