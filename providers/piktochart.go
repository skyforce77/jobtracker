package providers

type piktochart struct {
	greenhouse
}

// NewPiktochart returns a new provider
func NewPiktochart() Provider {
	return &piktochart{
		greenhouse{
			"Piktochart",
			"piktochart",
		},
	}
}
