package user_request_service

import (
	interfaces "arbyhunter/src/types/interfaces"

	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

type UserRequestService struct {
	server         *http.Server
	arbCoordinator interfaces.IArbCoordinator
}

func NewUserRequestService(arb_coordinator interfaces.IArbCoordinator) UserRequestService {
	// Start the server on port 8080
	port := os.Getenv("API_SERVER_PORT")

	server := &http.Server{
		Addr: ":" + port,
	}

	service := UserRequestService{
		server:         server,
		arbCoordinator: arb_coordinator,
	}

	// Routes
	mux := http.NewServeMux()
	mux.HandleFunc("/launchNodeAdaptor", service.launchNodeAdaptorHandler)
	mux.HandleFunc("/addPool", service.addPoolHandler)
	mux.HandleFunc("/healthCheck", service.healthCheckHandler)
	service.server.Handler = mux

	go func() {
		fmt.Printf("Starting API server on http://localhost:%s\n", port)
		// ListenAndServe is blocking
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Println("API server closed")
		} else if err != nil {
			fmt.Println("Error starting API server:", err)
		}
	}()

	return service
}

func CleanUpUserRequestService(service UserRequestService) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := service.server.Shutdown(ctx); err != nil {
		fmt.Println("API server Shutdown Failed:%+v\n", err)
	}
	fmt.Println("API server exited properly")
}
