package providers

type strait struct {
	lever
}

// NewStrait returns a new provider
func NewStrait() Provider {
	return &strait{
		lever{
			"Strait",
			"straitcapital",
		},
	}
}
