package providers

type anthropic struct {
	greenhouse
}

// NewAnthropic returns a new provider for Anthropic jobs
func NewAnthropic() Provider {
	return &anthropic{
		greenhouse{
			"Anthropic",
			"anthropic",
		},
	}
}
