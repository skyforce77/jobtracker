package providers

type giphy struct {
	greenhouse
}

// NewGiphy returns a new provider
func NewGiphy() Provider {
	return &giphy{
		greenhouse{
			"Giphy",
			"giphy",
		},
	}
}
