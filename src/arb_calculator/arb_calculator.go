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

	return &calculator
}
