package types

import (
	"arbyhunter/src/types/enums"
)

type LaunchNodeAdaptorDTO struct {
	Rawurl          string                `json:"raw_url"`
	NodeAdaptorType types.NodeAdaptorType `json:"node_adaptor_type"`
}
