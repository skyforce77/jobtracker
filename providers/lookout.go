package providers

type lookout struct {
	greenhouse
}

// NewLookout returns a new provider
func NewLookout() Provider {
	return &lookout{
		greenhouse{
			"Lookout",
			"lookout",
		},
	}
}
