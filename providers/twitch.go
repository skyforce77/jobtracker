package providers

type twitch struct {
	lever
}

func NewTwitch() *twitch {
	return &twitch{
		lever{
			"Twitch",
			"twitch",
		},
	}
}
