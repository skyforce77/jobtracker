package providers

type jaunt struct {
	greenhouse
}

// NewJauntVR returns a new provider
func NewJauntVR() Provider {
	return &jaunt{
		greenhouse{
			"JauntVR",
			"jaunt",
		},
	}
}
