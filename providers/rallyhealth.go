package providers

type rallyhealth struct {
	greenhouse
}

// NewRallyHealth returns a new provider
func NewRallyHealth() Provider {
	return &rallyhealth{
		greenhouse{
			"RallyHealth",
			"rallyhealth",
		},
	}
}
