package providers

type udemy struct {
	lever
}

// NewUdemy returns a new provider
func NewUdemy() Provider {
	return &udemy{
		lever{
			"Udemy",
			"udemy",
		},
	}
}
