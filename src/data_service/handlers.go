package data_service

import (
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (service *DataService) launchNodeAdaptorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: Invalid request method")
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
		fmt.Printf("launchNodeAdaptorHandler error: could not unmarshal body: %s \nerror: %s", body, err)
		return
	}

	fmt.Printf("launchNodeAdaptorHandler request: %+v", dto)

	// Input validation
	if dto.NodeAdaptorType < 0 || dto.NodeAdaptorType >= enums.MAX_VAL_NodeAdaptorType {
		http.Error(w, "invalid node_adaptor_type", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: invalid node_adaptor_type")
		return
	}

	// Send validated request to ArbCalculator

	response := service.arbCalculator.LaunchNodeAdaptor(dto)

	// Wait for response (with timeout) from ArbCalculator
	// TODO - timeout
	// timeout_timer := time.NewTimer(5 * time.Second)
	// defer timeout_timer.Stop()

	if response.Code != 200 {
		http.Error(w, "Failed to launch node adaptor: "+response.Message, http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: %+v", response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Launched Node Adaptor successfully"))
}

func (service *DataService) addPoolHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: Invalid request method")
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
		fmt.Printf("addPoolHandler error: could not unmarshal body: %s \nerror: %s", body, err)
		return
	}

	fmt.Printf("addPoolHandler request: %+v", dto)

	// Input validation
	if dto.NodeAdaptorType < 0 || dto.NodeAdaptorType >= enums.MAX_VAL_NodeAdaptorType {
		http.Error(w, "invalid node_adaptor_type", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: invalid node_adaptor_type")
		return
	}

	if dto.ProtocolAdaptorType < 0 || dto.ProtocolAdaptorType >= enums.MAX_VAL_ProtocolAdaptorType {
		http.Error(w, "invalid protocol_adaptor_type", http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: invalid protocol_adaptor_type")
		return
	}

	// TODO - Validate addresses in ProtocolAdaptor

	response := service.arbCalculator.AddPool(dto)

	// TODO do timeout
	// Wait for response (with timeout) from ArbCalculator
	// timeout_timer := time.NewTimer(5 * time.Second)
	// defer timeout_timer.Stop()

	if response.Code != 200 {
		http.Error(w, "Failed to add pool: "+response.Message, http.StatusBadRequest)
		fmt.Printf("addPoolHandler error: %+v", response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added pool successfully"))
}
