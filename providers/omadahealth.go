package providers

type omadahealth struct {
	greenhouse
}

// NewOmadaHealth returns a new provider
func NewOmadaHealth() Provider {
	return &omadahealth{
		greenhouse{
			"OmadaHealth",
			"omadahealth",
		},
	}
}
