package node_adaptor

import (
	dtos "arbyhunter/src/types/dtos"

	"github.com/ethereum/go-ethereum/ethclient"

	"fmt"
)

type NodeAdaptorEVM struct {
	blockchain_client *ethclient.Client
}

func NewNodeAdaptorEVM(rawurl string) *NodeAdaptorEVM {
	client, err := ethclient.Dial(rawurl)

	if err != nil {
		fmt.Println("NewNodeAdaptorEVM: Error connecting to", rawurl)
		return nil
	}

	fmt.Println("NewNodeAdaptorEVM success")
	return &NodeAdaptorEVM{
		blockchain_client: client,
	}
}

func (service *NodeAdaptorEVM) AddPool(dto dtos.AddPoolDTO) {
	println("AddPool")
}
