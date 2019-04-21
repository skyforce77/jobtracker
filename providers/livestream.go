package providers

type livestream struct {
	greenhouse
}

// NewLivestream returns a new provider
func NewLivestream() Provider {
	return &livestream{
		greenhouse{
			"Livestream",
			"livestream",
		},
	}
}
