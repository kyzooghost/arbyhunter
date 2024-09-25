package arb_calculator

import (
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"
	interfaces "arbyhunter/src/types/interfaces"
	models "arbyhunter/src/types/models"

	"fmt"
)

type ArbCalculator struct {
	dataServiceRequestChannel  chan *models.DataServiceRequest
	dataServiceResponseChannel chan *models.DataServiceResponse
	nodeAdaptors               map[enums.NodeAdaptorType]interfaces.INodeAdaptor
}

func NewArbCalculator(dataServiceRequestChannel chan *models.DataServiceRequest, dataServiceResponseChannel chan *models.DataServiceResponse) *ArbCalculator {
	fmt.Println("NewArbCalculator")
	calculator := ArbCalculator{
		dataServiceRequestChannel:  dataServiceRequestChannel,
		dataServiceResponseChannel: dataServiceResponseChannel,
		nodeAdaptors:               make(map[enums.NodeAdaptorType]interfaces.INodeAdaptor),
	}

	go func() {
		for {
			select {
			case request := <-dataServiceRequestChannel:
				var resp models.DataServiceResponse

				switch dto := request.Dto.(type) {
				default:
					resp.Code = 400
					resp.Message = "Unexpected DTO type"
					resp.Data = nil
					fmt.Printf("unexpected DTO type %T", dto)
				case dtos.LaunchNodeAdaptorDTO:
					resp = calculator.launchNodeAdaptorHandler(dto)
				case dtos.AddPoolDTO:
					resp = calculator.addPoolHandler(dto)
				}

				resp.RequestId = request.RequestId
				dataServiceResponseChannel <- &resp
			}
		}
	}()

	return &calculator
}
