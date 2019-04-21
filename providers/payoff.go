package providers

type payoff struct {
	greenhouse
}

// NewPayoff returns a new provider
func NewPayoff() Provider {
	return &payoff{
		greenhouse{
			"Payoff",
			"payoff",
		},
	}
}
