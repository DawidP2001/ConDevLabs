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
