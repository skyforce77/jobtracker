package providers

type doctrine struct {
	lever
}

// NewDoctrine returns a new provider
func NewDoctrine() Provider {
	return &doctrine{
		lever{
			"Doctrine",
			"doctrine",
		},
	}
}
