package main

import (
	"fmt"
	"time"
)

func main() {
	RequestQueue := make(chan Request, 5)

	for WorkerPool := 1; WorkerPool <= 3; WorkerPool++ {
		go ElevatorWorker(WorkerPool, RequestQueue)
	}

	for Burst := 1; Burst <= 10; Burst++ {
		req := Request{
			ID: Burst, Floor: Burst + 3, Type: "Normal"}

		fmt.Printf("Dispatcher: Sending Request %d to queue...\n", Burst)
		RequestQueue <- req

		fmt.Print("Queue Load: %d/%d\n", len(RequestQueue), cap(RequestQueue))

	}
	close(RequestQueue)

	time.Sleep(10 * time.Second) // Wait for workers to finish processing
	fmt.Println("Day 2 Lab Complete: All request processed.")
}
