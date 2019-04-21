package providers

type shazam struct {
	greenhouse
}

// NewShazam returns a new provider
func NewShazam() Provider {
	return &shazam{
		greenhouse{
			"Shazam",
			"shazam",
		},
	}
}
