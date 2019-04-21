package providers

type applovin struct {
	greenhouse
}

// NewAppLovin returns a new provider
func NewAppLovin() Provider {
	return &applovin{
		greenhouse{
			"AppLovin",
			"applovin",
		},
	}
}
