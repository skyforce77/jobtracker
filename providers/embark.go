package providers

type embark struct {
	greenhouse
}

// NewEmbark returns a new provider
func NewEmbark() Provider {
	return &embark{
		greenhouse{
			"Embark",
			"embark",
		},
	}
}
