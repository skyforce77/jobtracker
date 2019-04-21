package providers

type noom struct {
	greenhouse
}

// NewNoom returns a new provider
func NewNoom() Provider {
	return &noom{
		greenhouse{
			"Noom",
			"noom",
		},
	}
}
