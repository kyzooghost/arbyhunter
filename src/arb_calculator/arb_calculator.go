package arb_calculator

import (
	enums "arbyhunter/src/types/enums"
	interfaces "arbyhunter/src/types/interfaces"

	"fmt"
)

type ArbCalculator struct {
	nodeAdaptors map[enums.NodeAdaptorType]interfaces.INodeAdaptor
}

func NewArbCalculator() *ArbCalculator {
	fmt.Println("NewArbCalculator")
	calculator := ArbCalculator{
		nodeAdaptors: make(map[enums.NodeAdaptorType]interfaces.INodeAdaptor),
	}

	// go func() {
	// 	for {
	// 		select {
	// 		case request := <-dataServiceRequestChannel:
	// 			var resp models.DataServiceResponse

	// 			switch dto := request.Dto.(type) {
	// 			default:
	// 				resp.Code = 400
	// 				resp.Message = "Unexpected DTO type"
	// 				resp.Data = nil
	// 				fmt.Printf("unexpected DTO type %T", dto)
	// 			case dtos.LaunchNodeAdaptorDTO:
	// 				resp = calculator.launchNodeAdaptorHandler(dto)
	// 			case dtos.AddPoolDTO:
	// 				resp = calculator.addPoolHandler(dto)
	// 			}

	// 			resp.RequestId = request.RequestId
	// 			dataServiceResponseChannel <- &resp
	// 		}
	// 	}
	// }()

	return &calculator
}
