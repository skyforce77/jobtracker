package providers

type spredfast struct {
	greenhouse
}

// NewSpredfast returns a new provider
func NewSpredfast() Provider {
	return &spredfast{
		greenhouse{
			"Spredfast",
			"spredfast",
		},
	}
}
