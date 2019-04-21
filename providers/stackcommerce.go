package providers

type stackcommerce struct {
	greenhouse
}

// NewStackCommerce returns a new provider
func NewStackCommerce() Provider {
	return &stackcommerce{
		greenhouse{
			"StackCommerce",
			"stackcommerce",
		},
	}
}
