package main

import "time"

func main() {
	userChan := make(chan Request)
	adminChan := make(chan Request)

	go StartHandler(userChan, adminChan)

	go func ()  {
		userChan <- Request{ID: 1, Type: "USER", Content: "Requesting floor 5"}
		
	}()

	go func ()  {
		time.Sleep(2 * time.Second)
		adminChan <- Request{ID: 99, Type: "ADMIN", Content: "Emergency Brake Test Required"}
	}()

	time.Sleep(12 * time.Second)
}