package main

import (
	"fmt"
	"time"
)

type Status int

const (
	Idle   Status = 0
	GoUp   Status = 1
	GoDown Status = 2
)

type Request struct {
	currentFloor int
	destFloor    int
}

type Elevator struct {
	currentFloor int
	status       Status
}

func main() {

	elevator1 := Elevator{currentFloor: 1, status: 1}
	requests := []Request{{1, 2}, {2, 2}}
	elevator1.Process(requests)
}

func (e *Elevator) Process(requests []Request) {

	//Todo assign each request for each elevator
	for _, request := range requests {

		for request.currentFloor != e.currentFloor {
			fmt.Println("Process")
			if request.currentFloor > e.currentFloor {
				e.status = 1
				e.currentFloor++
				time.Sleep(2 * time.Second)
				fmt.Println("Elevator going up. Current floor is: ", e.currentFloor)
			}
			if request.currentFloor < e.currentFloor {
				e.status = 2
				e.currentFloor--
				time.Sleep(2 * time.Second)
				fmt.Println("Elevator going down. Current floor is: ", e.currentFloor)
			}
		}
	}

}
