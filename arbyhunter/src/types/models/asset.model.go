package types

type Asset struct {
	// Use as unique ID for asset, prefer Coingecko ticker, standardized to uppercase
	Ticker string `json:"ticker"`
	// Blockchain address of token
	Address string  `json:"address"`
	Name    *string `json:"name,omitempty"`
}
