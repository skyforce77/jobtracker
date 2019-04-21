package providers

type persado struct {
	greenhouse
}

// NewPersado returns a new provider
func NewPersado() Provider {
	return &persado{
		greenhouse{
			"Persado",
			"persado",
		},
	}
}
