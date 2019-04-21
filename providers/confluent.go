package providers

type confluent struct {
	lever
}

// NewConfluent returns a new provider
func NewConfluent() Provider {
	return &confluent{
		lever{
			"Confluent",
			"confluent",
		},
	}
}
