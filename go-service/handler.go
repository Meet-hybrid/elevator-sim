package main

import (
	"fmt"
	"time"
)

func ElevatorWorker(id int, queue <-chan Request) {
	for req := range queue {
		fmt.Printf("Elevator [%d]: Processing Request ID %d (Floor %d)\n", id, req.ID, req.Floor)

		time.Sleep(1 * time.Second) // Simulate processing time

		fmt.Printf("Elevator [%d]: Task ID %d Complete. \n", id, req.ID)
	}
}