package providers

type dubsmash struct {
	greenhouse
}

// NewDubsmash returns a new provider
func NewDubsmash() Provider {
	return &dubsmash{
		greenhouse{
			"Dubsmash",
			"dubsmash",
		},
	}
}
