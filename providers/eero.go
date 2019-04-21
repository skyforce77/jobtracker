package providers

type eero struct {
	greenhouse
}

// NewEero returns a new provider
func NewEero() Provider {
	return &eero{
		greenhouse{
			"Eero",
			"eero",
		},
	}
}
