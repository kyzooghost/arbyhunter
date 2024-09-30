package types

import (
	dtos "arbyhunter/src/types/dtos"
)

type INodeAdaptor interface {
	// Being lazy re: defining new types, want to reuse this AddPoolDTO type even though don't need 'NodeAdaptorType' at this point
	AddPool(dto dtos.AddPoolDTO)
}
