package types

type IArbCalculator interface {
	LaunchNodeAdaptor()
	AddPool()
	ScanArbitrages()
}
