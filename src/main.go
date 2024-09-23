package main

import (
	"github.com/joho/godotenv"

	"arbyhunter/src/arb_calculator"
	"arbyhunter/src/data_service"

	"fmt"
	"log"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Hello, Go!")
	apiCall_to_arbCalculator_channel := make(chan interface{})
	go data_service.InitDataService(apiCall_to_arbCalculator_channel)
	go arb_calculator.InitArbCalculator(apiCall_to_arbCalculator_channel)

	// Infinite loop
	// for {
	// Throttle for loop to lower CPU usage
	// }
	for {
		time.Sleep(100)

		// select {
		// case msg := <-dataService_to_arbCalculator_channel:
		// 	fmt.Println(msg)
		// }
	}
}
