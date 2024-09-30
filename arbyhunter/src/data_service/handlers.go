package data_service

import (
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"

	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (service *DataService) launchNodeAdaptorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: Invalid request method\n")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: could not read request body: %s\n", err)
		return
	}

	var dto dtos.LaunchNodeAdaptorDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: could not unmarshal body: %s \nerror: %s\n", body, err)
		return
	}

	fmt.Printf("launchNodeAdaptorHandler request: %+v\n", dto)

	// Input validation
	if dto.NodeAdaptorType < 0 || dto.NodeAdaptorType >= enums.MAX_VAL_NodeAdaptorType {
		http.Error(w, "invalid node_adaptor_type", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: invalid node_adaptor_type\n")
		return
	}

	// Send validated request to ArbCalculator

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	response := service.arbCalculator.LaunchNodeAdaptor(ctx, dto)

	if response.Code != 200 {
		http.Error(w, "Failed to launch node adaptor: "+response.Message, http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: %+v\n", response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Launched Node Adaptor successfully"))
}

func (service *DataService) addPoolHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: Invalid request method\n")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: could not read request body: %s\n", err)
		return
	}

	var dto dtos.AddPoolDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: could not unmarshal body: %s \nerror: %s\n", body, err)
		return
	}

	fmt.Printf("addPoolHandler request: %+v\n", dto)

	// Input validation
	if dto.NodeAdaptorType < 0 || dto.NodeAdaptorType >= enums.MAX_VAL_NodeAdaptorType {
		http.Error(w, "invalid node_adaptor_type", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: invalid node_adaptor_type\n")
		return
	}

	if dto.ProtocolAdaptorType < 0 || dto.ProtocolAdaptorType >= enums.MAX_VAL_ProtocolAdaptorType {
		http.Error(w, "invalid protocol_adaptor_type", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: invalid protocol_adaptor_type\n")
		return
	}

	// TODO - Validate addresses in ProtocolAdaptor

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	response := service.arbCalculator.AddPool(ctx, dto)

	if response.Code != 200 {
		http.Error(w, "Failed to add pool: "+response.Message, http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: %+v\n", response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added pool successfully"))
}

func (service *DataService) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		fmt.Printf("healthCheckHandler error: Invalid request method\n")
		return
	}

	fmt.Printf("healthCheckHandler request\n")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	response := service.arbCalculator.HealthCheck(ctx)

	if response.Code != 200 {
		http.Error(w, "Failed health check: "+response.Message, http.StatusBadRequest)
		fmt.Printf("healthCheckHandler error: %+v\n", response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response.Message))
}
