package providers

type braintree struct {
	greenhouse
}

// NewBraintree returns a new provider
func NewBraintree() Provider {
	return &braintree{
		greenhouse{
			"Braintree",
			"braintree",
		},
	}
}
