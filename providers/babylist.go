package providers

type babylist struct {
	lever
}

// NewBabylist returns a new provider
func NewBabylist() Provider {
	return &babylist{
		lever{
			"Babylist",
			"babylist",
		},
	}
}
