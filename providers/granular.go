package providers

type granular struct {
	greenhouse
}

// NewGranular returns a new provider
func NewGranular() Provider {
	return &granular{
		greenhouse{
			"Granular",
			"granular",
		},
	}
}
