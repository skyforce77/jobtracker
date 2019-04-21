package providers

type roundsphere struct {
	greenhouse
}

// NewRoundSphere returns a new provider
func NewRoundSphere() Provider {
	return &roundsphere{
		greenhouse{
			"RoundSphere",
			"roundsphere",
		},
	}
}
