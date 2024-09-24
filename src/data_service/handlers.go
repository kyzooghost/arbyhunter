package data_service

import (
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"
	models "arbyhunter/src/types/models"

	"github.com/google/uuid"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func (service *DataService) launchNodeAdaptorHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invaild request", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: could not read request body: %s\n", err)
	}

	var dto dtos.LaunchNodeAdaptorDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: could not unmarshal body: %s \nerror: %s", body, err)
		return
	}

	fmt.Printf("launchNodeAdaptorHandler request: raw_url=%s, node_adaptor_type=%d\n", dto.Rawurl, dto.NodeAdaptorType)

	// Input validation
	if dto.NodeAdaptorType < 0 || dto.NodeAdaptorType >= enums.MAX_NodeAdaptorType_VAL {
		http.Error(w, "invalid node_adaptor_type", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: invalid node_adaptor_type")
		return
	}

	// Send validated request to ArbCalculator
	request := models.DataServiceRequest{
		RequestId: uuid.New().String(),
		Dto:       dto,
	}

	service.dataServiceRequestChannel <- &request

	// Wait for response (with timeout) from ArbCalculator
	timeout_timer := time.NewTimer(5 * time.Second)
	defer timeout_timer.Stop()
	isResponseReceived := false
	var response *models.DataServiceResponse

	for !isResponseReceived {
		select {
		case response = <-service.dataServiceResponseChannel:
			isResponseReceived = response.RequestId == request.RequestId
		case <-timeout_timer.C:
			http.Error(w, "Request timeout", http.StatusBadRequest)
			fmt.Printf("launchNodeAdaptorHandler error: timed out after sending request to ArbCalculator")
			return
		}
	}

	if response.Code != 200 {
		http.Error(w, "Failed to launch node adaptor: "+response.Message, http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: %+v", response)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Launched Node Adaptor successfully"))
}
