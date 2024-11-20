package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func simulateRegistration(userID int, wg *sync.WaitGroup) {
	// This will only decrement if wg is not nil (applicable to sequential mode)
	defer func() {
		if wg != nil {
			wg.Done()
		}
	}()
	
	// userAuthServiceUrl1 := "http://localhost:8080"
	userAuthServiceUrl2 := "http://34.28.222.86"

	// Simulate user registration
	email := fmt.Sprintf("user%d@example.com", userID)
	payload := []byte(fmt.Sprintf(`{"email": "%s", "password": "securepassword", "name": "User %d"}`, email, userID))

	resp, err := http.Post(userAuthServiceUrl2+"/registration", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Failed to register user %d: %v\n", userID, err)
		return
	}
	defer resp.Body.Close()

	log.Printf("User %d registered with status: %d\n", userID, resp.StatusCode)
}

func runSequentialSimulation(numUsers int, delay time.Duration) {
	log.Println("Starting sequential simulation...")
	start := time.Now()
	for i := 1; i <= numUsers; i++ {
		simulateRegistration(i, nil) // nil because this is sequential
		time.Sleep(delay)
	}
	log.Printf("Sequential simulation completed in %v\n", time.Since(start))
}

func runConcurrentSimulation(startUserID, numUsers int, delay time.Duration) {
	log.Println("Starting concurrent simulation...")
	start := time.Now()
	var wg sync.WaitGroup
	for i := startUserID; i < startUserID+numUsers; i++ {
		wg.Add(1)
		go simulateRegistration(i, &wg)
		time.Sleep(delay)
	}
	wg.Wait()
	log.Printf("Concurrent simulation completed in %v\n", time.Since(start))
}

func main() {
	const numUsersPerMode = 60
	const delayBetweenRequests = 200 * time.Millisecond

	// run sequential simulation
	runSequentialSimulation(numUsersPerMode, delayBetweenRequests)

	// run concurrent siumulation
	runConcurrentSimulation(numUsersPerMode+1, numUsersPerMode, delayBetweenRequests)

}
