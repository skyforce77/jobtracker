package providers

type splash struct {
	greenhouse
}

// NewSplash returns a new provider
func NewSplash() Provider {
	return &splash{
		greenhouse{
			"Splash",
			"splash",
		},
	}
}
