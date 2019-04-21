package providers

type kickstarter struct {
	lever
}

// NewKickStarter returns a new provider
func NewKickStarter() Provider {
	return &kickstarter{
		lever{
			"KickStarter",
			"kickstarter",
		},
	}
}
