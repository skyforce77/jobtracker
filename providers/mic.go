package providers

type mic struct {
	greenhouse
}

// NewMic returns a new provider
func NewMic() Provider {
	return &mic{
		greenhouse{
			"Mic",
			"mic",
		},
	}
}
