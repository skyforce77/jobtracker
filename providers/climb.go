package providers

type climb struct {
	greenhouse
}

// NewClimb returns a new provider
func NewClimb() Provider {
	return &climb{
		greenhouse{
			"Climb",
			"climb",
		},
	}
}
