package providers

type stripe struct {
	greenhouse
}

// NewStripe returns a new provider for Stripe jobs
func NewStripe() Provider {
	return &stripe{
		greenhouse{
			"Stripe",
			"stripe",
		},
	}
}
