package providers

type pindropsecurity struct {
	greenhouse
}

// NewPindrop returns a new provider
func NewPindrop() Provider {
	return &pindropsecurity{
		greenhouse{
			"Pindrop",
			"pindropsecurity",
		},
	}
}
