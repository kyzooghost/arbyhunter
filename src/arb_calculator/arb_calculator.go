package arb_calculator

import (
	models "arbyhunter/src/types/models"

	"fmt"
	"time"
)

func InitArbCalculator(dataServiceRequestChannel chan *models.DataServiceRequest, dataServiceResponseChannel chan *models.DataServiceResponse) {
	fmt.Println("InitArbCalculator")

	for {
		select {
		case request := <-dataServiceRequestChannel:
			time.Sleep(1 * time.Second)
			resp := models.DataServiceResponse{
				RequestId: request.RequestId,
				Data:      nil,
			}
			dataServiceResponseChannel <- &resp
		}
	}
}
