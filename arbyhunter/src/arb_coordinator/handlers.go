package arb_coordinator

import (
	"arbyhunter/src/node_adaptor"
	dtos "arbyhunter/src/types/dtos"
	enums "arbyhunter/src/types/enums"
	models "arbyhunter/src/types/models"

	zmq "github.com/go-zeromq/zmq4"

	"context"
	"fmt"
	"os"
	"time"
)

func (service ArbCoordinator) LaunchNodeAdaptor(ctx context.Context, dto dtos.LaunchNodeAdaptorDTO) models.UserResponse {
	fmt.Printf("LaunchNodeAdaptor started\n")
	switch nodeAdaptorType := dto.NodeAdaptorType; nodeAdaptorType {
	case enums.EVM:

		// Timeout protection
		adaptor, err := node_adaptor.NewNodeAdaptorEVM(ctx, dto.Rawurl)
		if err != nil {
			return models.UserResponse{
				Code:    400,
				Message: err.Error(),
				Data:    nil,
			}
		}
		if adaptor == nil {
			return models.UserResponse{
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
	return models.UserResponse{
		Code:    200,
		Message: "",
		Data:    nil,
	}
}

func (service ArbCoordinator) AddPool(ctx context.Context, dto dtos.AddPoolDTO) models.UserResponse {
	fmt.Printf("AddPool started\n")
	nodeAdaptor, exists := service.nodeAdaptors[dto.NodeAdaptorType]
	if !exists {
		return models.UserResponse{
			Code:    400,
			Message: fmt.Sprintf("NodeAdaptor type %d not yet launched", dto.NodeAdaptorType),
			Data:    nil,
		}
	}

	// TODO Check for success/failure resp
	nodeAdaptor.AddPool(dto)

	fmt.Printf("AddPool succeeded\n")
	return models.UserResponse{
		Code:    200,
		Message: "",
		Data:    nil,
	}
}

// Send 'hello' string to the ArbyScanner server
func (service ArbCoordinator) HealthCheck(ctx context.Context) models.UserResponse {
	fmt.Println("HealthCheck started\n")

	socket, err := service.connectToArbScanner(ctx)
	if err != nil {
		fmt.Println("arb_coordinator.HealthCheck - failed to connect to ArbScanner server: %w", err)
		return models.UserResponse{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}
	}
	defer socket.Close()

	// Send msg
	msg := zmq.NewMsgString("hello")
	fmt.Println("Sending message to ArbScanner server: ", msg)
	if err := socket.Send(msg); err != nil {
		fmt.Println("arb_coordinator.HealthCheck - failed send message to ArbScanner server: %w", err)
		return models.UserResponse{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}
	}

	// Wait for reply.
	r, err := socket.Recv()
	if err != nil {
		fmt.Println("arb_coordinator.HealthCheck - failed receive message from ArbScanner server: %w", err)
		return models.UserResponse{
			Code:    500,
			Message: err.Error(),
			Data:    nil,
		}
	}
	fmt.Println("received message from ArbScanner server: ", r.String())

	fmt.Println("HealthCheck succeeded\n")
	return models.UserResponse{
		Code:    200,
		Message: "HealthCheck success\n",
		Data:    nil,
	}
}

// Return unique Socket instance, rather than reuse single Socket instance as we are advised 'do not use the same socket from multiple threads'
func (service ArbCoordinator) connectToArbScanner(ctx context.Context) (zmq.Socket, error) {
	socket := zmq.NewReq(ctx, zmq.WithDialerRetry(time.Second))
	ipcEndpoint := os.Getenv("ARB_SCANNER_IPC_ENDPOINT")
	fmt.Printf("Connecting to ArbScanner server at %s\n", ipcEndpoint)
	if err := socket.Dial(ipcEndpoint); err != nil {
		fmt.Printf("Error connecting to ArbScanner server: %w\n", err)
		return nil, err
	}
	fmt.Println("Successfully connected to ArbScanner server")
	return socket, nil
}
