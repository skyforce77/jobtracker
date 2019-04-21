package providers

type hottopic struct {
	lever
}

// NewHottopic returns a new provider
func NewHottopic() Provider {
	return &hottopic{
		lever{
			"Hot Topic",
			"hottopic",
		},
	}
}
