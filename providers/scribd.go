package providers

type scribd struct {
	lever
}

// NewScribd returns a new provider
func NewScribd() Provider {
	return &scribd{
		lever{
			"Scribd",
			"scribd",
		},
	}
}
