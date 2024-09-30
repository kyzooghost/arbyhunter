package node_adaptor

import (
	dtos "arbyhunter/src/types/dtos"

	"github.com/ethereum/go-ethereum/ethclient"

	"context"
	"fmt"
)

type NodeAdaptorEVM struct {
	blockchain_client *ethclient.Client
}

func NewNodeAdaptorEVM(ctx context.Context, rawurl string) (*NodeAdaptorEVM, error) {
	resultChan := make(chan *ethclient.Client)
	errChan := make(chan error)
	go func() {
		client, err := ethclient.Dial(rawurl)
		if err != nil {
			errChan <- err
		}
		resultChan <- client
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("NewNodeAdaptorEVM request timed out")
	case err := <-errChan:
		errString := fmt.Sprintf("NewNodeAdaptorEVM: Error connecting to"+rawurl+": %s", err.Error()+"\n")
		fmt.Println(errString)
		return nil, fmt.Errorf(errString)
	case client := <-resultChan:
		fmt.Println("NewNodeAdaptorEVM success")
		return &NodeAdaptorEVM{
			blockchain_client: client,
		}, nil
	}
}

func (service *NodeAdaptorEVM) AddPool(dto dtos.AddPoolDTO) {
	println("AddPool")
}
