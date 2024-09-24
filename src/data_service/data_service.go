package data_service

import (
	"context"
	"errors"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/launchNodeAdaptor", launchNodeAdaptorHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	go func() {
		fmt.Println("Starting API server on port", port)
		// ListenAndServe is blocking
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("API server closed")
		} else if err != nil {
			fmt.Println("Error starting API server:", err)
		}
	}()

	return &DataService{
		server: server,
	}
}

func CleanUpDataService(service *DataService) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := service.server.Shutdown(ctx); err != nil {
		fmt.Printf("API server Shutdown Failed:%+v", err)
	}
	fmt.Println("API server exited properly")
}
