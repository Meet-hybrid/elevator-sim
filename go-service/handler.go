package main

import (
	"fmt"
	"time"
)

func StartHandler(userRequests <-chan Request, adminRequests <-chan Request) {
	for {
		select {
		case req := <-adminRequests:
			fmt.Printf("ADMIN [ID%d]: %s\n", req.ID, req.Content)
		case req := <-userRequests:
			fmt.Printf("USER [ID%d]: %s\n", req.ID, req.Content)
		case <-time.After(10 * time.Second):
			fmt.Println("No messages received in the last 5 seconds.")
		}
	}
}