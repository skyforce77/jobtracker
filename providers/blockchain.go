package providers

type blockchain struct {
	greenhouse
}

// NewBlockchain returns a new provider
func NewBlockchain() Provider {
	return &blockchain{
		greenhouse{
			"Blockchain",
			"blockchain",
		},
	}
}
