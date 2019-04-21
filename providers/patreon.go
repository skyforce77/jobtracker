package providers

type patreon struct {
	greenhouse
}

// NewPatreon returns a new provider
func NewPatreon() Provider {
	return &patreon{
		greenhouse{
			"Patreon",
			"patreon",
		},
	}
}
