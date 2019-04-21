package providers

type mixpanel struct {
	greenhouse
}

// NewMixpanel returns a new provider
func NewMixpanel() Provider {
	return &mixpanel{
		greenhouse{
			"Mixpanel",
			"mixpanel",
		},
	}
}
