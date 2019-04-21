package providers

type outreach struct {
	lever
}

// NewOutreach returns a new provider
func NewOutreach() Provider {
	return &outreach{
		lever{
			"Outreach",
			"outreach",
		},
	}
}
