package types

import (
	enums "arbyhunter/src/types/enums"
	models "arbyhunter/src/types/models"
)

type AddPoolDTO struct {
	NodeAdaptorType     enums.NodeAdaptorType     `json:"node_adaptor_type"`
	ProtocolAdaptorType enums.ProtocolAdaptorType `json:"protocol_adaptor_type"`
	PoolAddress         string                    `json:"pool_address"`
	Assets              []models.Asset            `json:"assets"`
}
