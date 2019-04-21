package providers

type truemotion struct {
	greenhouse
}

// NewTrueMotion returns a new provider
func NewTrueMotion() Provider {
	return &truemotion{
		greenhouse{
			"TrueMotion",
			"truemotion",
		},
	}
}
