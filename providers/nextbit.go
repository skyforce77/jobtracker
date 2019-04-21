package providers

type nextbit struct {
	greenhouse
}

// NewNextbit returns a new provider
func NewNextbit() Provider {
	return &nextbit{
		greenhouse{
			"Nextbit",
			"nextbit",
		},
	}
}
