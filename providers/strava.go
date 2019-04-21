package providers

type strava struct {
	greenhouse
}

// NewStrava returns a new provider
func NewStrava() Provider {
	return &strava{
		greenhouse{
			"Strava",
			"strava",
		},
	}
}
