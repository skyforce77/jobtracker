package providers

type tripadvisor struct {
	greenhouse
}

// NewTripAdvisor returns a new provider
func NewTripAdvisor() Provider {
	return &tripadvisor{
		greenhouse{
			"TripAdvisor",
			"tripadvisor",
		},
	}
}
