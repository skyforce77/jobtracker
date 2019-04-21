package providers

type carousell struct {
	greenhouse
}

// NewCarousell returns a new provider
func NewCarousell() Provider {
	return &carousell{
		greenhouse{
			"Carousell",
			"carousell",
		},
	}
}
