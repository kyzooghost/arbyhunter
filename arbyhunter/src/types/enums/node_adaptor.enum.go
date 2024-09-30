package types

type NodeAdaptorType int

// Technique to validate enum values - https://stackoverflow.com/a/37502382
const (
	EVM NodeAdaptorType = iota
	SOLANA
	MAX_VAL_NodeAdaptorType
)
