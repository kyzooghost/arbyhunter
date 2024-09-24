package arb_calculator

import (
	"arbyhunter/src/node_adaptor"
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"
	models "arbyhunter/src/types/models"

	"fmt"
)

func (service *ArbCalculator) launchNodeAdaptorHandler(dto dtos.LaunchNodeAdaptorDTO) models.DataServiceResponse {
	switch nodeAdaptorType := dto.NodeAdaptorType; nodeAdaptorType {
	case enums.EVM:
		node_adaptor := node_adaptor.NewNodeAdaptorEVM(dto.Rawurl)
		if node_adaptor == nil {
			return models.DataServiceResponse{
				Code:    400,
				Message: "NewNodeAdaptorEVM failed",
				Data:    nil,
			}
		}
		service.nodeAdaptors[enums.EVM] = node_adaptor
	case enums.SOLANA:
		fmt.Printf("launchNodeAdaptorSolana")
	}

	return models.DataServiceResponse{
		Code:    200,
		Message: "",
		Data:    nil,
	}
}
