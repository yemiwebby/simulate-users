package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func simulateRegistration(userID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate user registration
	email := fmt.Sprintf("user%d@example.com", userID)
	payload := []byte(fmt.Sprintf(`{"email": "%s", "password": "securepassword", "name": "User %d"}`, email, userID))

	resp, err := http.Post("http://localhost:8080/registration", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Failed to register user %d: %v\n", userID, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("User %d registered with status: %d\n", userID, resp.StatusCode)
}

func main() {
	const numUsers = 50
	// const delayBetweenRequests = 200 * time.Millisecond
	var wg sync.WaitGroup

	for i := 1; i <= numUsers; i++ {
		wg.Add(1)
		go simulateRegistration(i, &wg)

		// time.Sleep(delayBetweenRequests)
	}

	wg.Wait()
	log.Println("Simulation complete: All users registered.")
}
