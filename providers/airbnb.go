package providers

type airbnb struct {
	greenhouse
}

// NewAirbnb returns a new provider
func NewAirbnb() Provider {
	return &airbnb{
		greenhouse{
			"Airbnb",
			"airbnb",
		},
	}
}
