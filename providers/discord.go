package providers

type discord struct {
	greenhouse
}

// NewDiscord returns a new provider for Discord jobs
func NewDiscord() Provider {
	return &discord{
		greenhouse{
			"Discord",
			"discord",
		},
	}
}
