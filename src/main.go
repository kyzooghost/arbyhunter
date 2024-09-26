package main

import (
	"github.com/joho/godotenv"

	"arbyhunter/src/arb_calculator"
	"arbyhunter/src/data_service"

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

	arb_calculator_instance := arb_calculator.NewArbCalculator()
	data_service_instance := data_service.NewDataService(arb_calculator_instance)

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
				data_service.CleanUpDataService(data_service_instance)
			}()
			wg.Wait()
			fmt.Println("Shutting down arbyhunter")
			return
		}
	}
}
