package providers

type spotify struct {
	lever
}

// NewSpotify returns a new provider for Spotify jobs
func NewSpotify() Provider {
	return &spotify{
		lever{
			"Spotify",
			"spotify",
		},
	}
}
