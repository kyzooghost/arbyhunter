package types

import (
	enums "arbyhunter/src/types/enums"
)

type LaunchNodeAdaptorDTO struct {
	Rawurl          string                `json:"raw_url"`
	NodeAdaptorType enums.NodeAdaptorType `json:"node_adaptor_type"`
}
