package main

import (
	"github.com/joho/godotenv"

	"arbyhunter/src/arb_calculator"
	"arbyhunter/src/data_service"
	models "arbyhunter/src/types/models"

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

	fmt.Println("Hello, Go!")
	dataServiceRequestChannel := make(chan *models.DataServiceRequest)
	dataServiceResponseChannel := make(chan *models.DataServiceResponse)

	data_service_instance := data_service.NewDataService(dataServiceRequestChannel, dataServiceResponseChannel)
	go arb_calculator.InitArbCalculator(dataServiceRequestChannel, dataServiceResponseChannel)

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
