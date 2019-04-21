package providers

type gusto struct {
	greenhouse
}

// NewGusto returns a new provider
func NewGusto() Provider {
	return &gusto{
		greenhouse{
			"Gusto",
			"gusto",
		},
	}
}
