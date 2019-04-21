package providers

type wework struct {
	greenhouse
}

// NewWeWork returns a new provider
func NewWeWork() Provider {
	return &wework{
		greenhouse{
			"WeWork",
			"wework",
		},
	}
}
