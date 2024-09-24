package data_service

import (
	"arbyhunter/src/types/dtos"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func launchNodeAdaptorHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "invaild request", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: could not read request body: %s\n", err)
	}

	var dto types.LaunchNodeAdaptorDTO
	err = json.Unmarshal(body, &dto)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		fmt.Printf("launchNodeAdaptorHandler error: could not unmarshal body: %s \nerror: %s", body, err)
		return
	}

	fmt.Printf("launchNodeAdaptorHandler request: raw_url=%s, node_adaptor_type=%s\n", dto.Rawurl, dto.NodeAdaptorType)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Launched Node Adaptor successfully"))
}
