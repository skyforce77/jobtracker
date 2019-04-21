package providers

type helix struct {
	greenhouse
}

// NewHelix returns a new provider
func NewHelix() Provider {
	return &helix{
		greenhouse{
			"Helix",
			"helix",
		},
	}
}
