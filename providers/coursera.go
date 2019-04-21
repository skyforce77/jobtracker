package providers

type coursera struct {
	lever
}

// NewCoursera returns a new provider
func NewCoursera() Provider {
	return &coursera{
		lever{
			"Coursera",
			"coursera",
		},
	}
}
