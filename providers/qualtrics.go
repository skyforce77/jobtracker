package providers

type qualtrics struct {
	greenhouse
}

// NewQualtrics returns a new provider
func NewQualtrics() Provider {
	return &qualtrics{
		greenhouse{
			"Qualtrics",
			"qualtrics",
		},
	}
}
