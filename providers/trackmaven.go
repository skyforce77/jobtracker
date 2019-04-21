package providers

type trackmaven struct {
	greenhouse
}

// NewTrackMaven returns a new provider
func NewTrackMaven() Provider {
	return &trackmaven{
		greenhouse{
			"TrackMaven",
			"trackmaven",
		},
	}
}
