package providers

type coinbase struct {
	greenhouse
}

// NewCoinbase returns a new provider for Coinbase jobs
func NewCoinbase() Provider {
	return &coinbase{
		greenhouse{
			"Coinbase",
			"coinbase",
		},
	}
}
