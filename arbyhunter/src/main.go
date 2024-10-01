package main

import (
	"github.com/joho/godotenv"

	"arbyhunter/src/arb_coordinator"
	"arbyhunter/src/user_request_service"

	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	arb_coordinator_instance := arb_coordinator.NewArbCoordinator()
	user_request_service_instance := user_request_service.NewUserRequestService(arb_coordinator_instance)

	// Cleanup
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-stop:
			fmt.Println("Cleaning up arbyhunter")
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				user_request_service.CleanUpUserRequestService(user_request_service_instance)
			}()
			wg.Wait()
			fmt.Println("Shutting down arbyhunter")
			return
		}
	}
}
