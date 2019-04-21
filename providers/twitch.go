package providers

type twitch struct {
	lever
}

// NewTwitch returns a new provider
func NewTwitch() Provider {
	return &twitch{
		lever{
			"Twitch",
			"twitch",
		},
	}
}
