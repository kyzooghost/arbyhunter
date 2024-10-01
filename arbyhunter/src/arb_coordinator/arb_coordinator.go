package arb_coordinator

import (
	enums "arbyhunter/src/types/enums"
	interfaces "arbyhunter/src/types/interfaces"

	"fmt"
)

type ArbCoordinator struct {
	nodeAdaptors map[enums.NodeAdaptorType]interfaces.INodeAdaptor
}

func NewArbCoordinator() *ArbCoordinator {
	fmt.Println("NewArbCoordinator")
	co := ArbCoordinator{
		nodeAdaptors: make(map[enums.NodeAdaptorType]interfaces.INodeAdaptor),
	}

	return &co
}
