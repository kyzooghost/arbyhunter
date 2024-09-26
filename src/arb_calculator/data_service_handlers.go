package arb_calculator

import (
	"arbyhunter/src/node_adaptor"
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"
	models "arbyhunter/src/types/models"

	"context"
	"fmt"
)

func (service *ArbCalculator) LaunchNodeAdaptor(ctx context.Context, dto dtos.LaunchNodeAdaptorDTO) models.DataServiceResponse {
	fmt.Printf("LaunchNodeAdaptor started\n")
	switch nodeAdaptorType := dto.NodeAdaptorType; nodeAdaptorType {
	case enums.EVM:

		// Timeout protection
		adaptor, err := node_adaptor.NewNodeAdaptorEVM(ctx, dto.Rawurl)
		if err != nil {
			return models.DataServiceResponse{
				Code:    400,
				Message: err.Error(),
				Data:    nil,
			}
		}
		if adaptor == nil {
			return models.DataServiceResponse{
				Code:    400,
				Message: "NewNodeAdaptorEVM failed",
				Data:    nil,
			}
		}
		service.nodeAdaptors[enums.EVM] = adaptor

	case enums.SOLANA:
		fmt.Printf("launchNodeAdaptorSolana\n")
	}

	fmt.Printf("LaunchNodeAdaptor succeeded\n")
	return models.DataServiceResponse{
		Code:    200,
		Message: "",
		Data:    nil,
	}
}

func (service *ArbCalculator) AddPool(ctx context.Context, dto dtos.AddPoolDTO) models.DataServiceResponse {
	fmt.Printf("AddPool started\n")
	nodeAdaptor, exists := service.nodeAdaptors[dto.NodeAdaptorType]
	if !exists {
		return models.DataServiceResponse{
			Code:    400,
			Message: fmt.Sprintf("NodeAdaptor type %d not yet launched", dto.NodeAdaptorType),
			Data:    nil,
		}
	}

	// TODO Check for success/failure resp
	nodeAdaptor.AddPool(dto)

	fmt.Printf("AddPool succeeded\n")
	return models.DataServiceResponse{
		Code:    200,
		Message: "",
		Data:    nil,
	}
}
