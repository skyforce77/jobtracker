package providers

type frankly struct {
	greenhouse
}

// NewFrankly returns a new provider
func NewFrankly() Provider {
	return &frankly{
		greenhouse{
			"Frankly",
			"frankly",
		},
	}
}
