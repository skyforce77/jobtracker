package providers

type britandco struct {
	greenhouse
}

// NewBritAndCo returns a new provider
func NewBritAndCo() Provider {
	return &britandco{
		greenhouse{
			"BritAndCo",
			"britandco",
		},
	}
}
