package providers

type bombfell struct {
	greenhouse
}

// NewBombfell returns a new provider
func NewBombfell() Provider {
	return &bombfell{
		greenhouse{
			"Bombfell",
			"bombfell",
		},
	}
}
