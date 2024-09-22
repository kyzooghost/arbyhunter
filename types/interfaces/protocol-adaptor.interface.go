package types

type IProtocolAdaptor interface {
	GetPoolPrices()
	ParseFilteredTx()
}
