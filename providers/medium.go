package providers

type medium struct {
	lever
}

// NewMedium returns a new provider
func NewMedium() Provider {
	return &medium{
		lever{
			"Medium",
			"medium",
		},
	}
}
