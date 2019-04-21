package providers

type envato struct {
	greenhouse
}

// NewEnvato returns a new provider
func NewEnvato() Provider {
	return &envato{
		greenhouse{
			"Envato",
			"envato",
		},
	}
}
