package providers

type journy struct {
	lever
}

// NewJourny returns a new provider
func NewJourny() Provider {
	return &journy{
		lever{
			"Journy",
			"gojourny",
		},
	}
}
