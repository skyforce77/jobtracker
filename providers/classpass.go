package providers

type classpass struct {
	greenhouse
}

// NewClassPass returns a new provider
func NewClassPass() Provider {
	return &classpass{
		greenhouse{
			"ClassPass",
			"classpass",
		},
	}
}
