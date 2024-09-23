package data_service

import (
	// "context"
	"fmt"
	"net/http"
	"os"
	// "os/signal"
	"sync"
	// "syscall"
	// "time"
)

func InitDataService(apiCall_to_arbCalculator_channel chan interface{}) {
	fmt.Println("InitDataService")
	startAPIServer()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func startAPIServer() {
	// Start the server on port 8080
	port := os.Getenv("API_SERVER_PORT")
	fmt.Println("Starting API server on port", port)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(helloHandler),
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// ListenAndServe is blocking
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Error starting API server:", err)
		}
	}()

	// Cleanup
	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// <-stop
	// fmt.Println("Shutting down API server")

	// // Create a context with a timeout for graceful shutdown
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// // defer cancel()
	// // // Shutdown the server gracefully
	// if err := server.Shutdown(ctx); err != nil {
	// 	fmt.Printf("Server Shutdown Failed:%+v", err)
	// }
	// wg.Wait()
	// fmt.Println("Server exited properly")
	// cancel()
}
