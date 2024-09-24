package data_service

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type DataService struct {
	server *http.Server
}

func NewDataService() *DataService {
	// Start the server on port 8080
	port := os.Getenv("API_SERVER_PORT")

	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(handler),
	}

	go func() {
		fmt.Println("Starting API server on port", port)
		// ListenAndServe is blocking
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("Error starting API server:", err)
		}
	}()

	return &DataService{
		server: server,
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func CleanUpDataService(service *DataService) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := service.server.Shutdown(ctx); err != nil {
		fmt.Printf("API server Shutdown Failed:%+v", err)
	}
	fmt.Println("API server exited properly")
}
