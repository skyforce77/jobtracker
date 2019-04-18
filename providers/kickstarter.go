package providers

type kickstarter struct {
	lever
}

func NewKickStarter() *kickstarter {
	return &kickstarter{
		lever{
			"KickStarter",
			"kickstarter",
		},
	}
}
