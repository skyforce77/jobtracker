package providers

type netskope struct {
	greenhouse
}

// NewNetskope returns a new provider
func NewNetskope() Provider {
	return &netskope{
		greenhouse{
			"Netskope",
			"netskope",
		},
	}
}
