package providers

type tremorvideo struct {
	greenhouse
}

// NewTremorVideo returns a new provider
func NewTremorVideo() Provider {
	return &tremorvideo{
		greenhouse{
			"TremorVideo",
			"tremorvideo",
		},
	}
}
