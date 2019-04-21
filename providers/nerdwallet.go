package providers

type nerdwallet struct {
	greenhouse
}

// NewNerdwallet returns a new provider
func NewNerdwallet() Provider {
	return &nerdwallet{
		greenhouse{
			"Nerdwallet",
			"nerdwallet",
		},
	}
}
