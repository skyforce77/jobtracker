package providers

type uber struct {
	greenhouse
}

// NewUber returns a new provider
func NewUber() Provider {
	return &uber{
		greenhouse{
			"Uber",
			"uber",
		},
	}
}
