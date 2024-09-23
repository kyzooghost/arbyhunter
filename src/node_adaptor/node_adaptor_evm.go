package node_adaptor

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
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

	return &NodeAdaptorEVM{
		blockchain_client: client,
	}
}
