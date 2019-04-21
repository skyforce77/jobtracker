package providers

type quantifind struct {
	greenhouse
}

// NewQuantifind returns a new provider
func NewQuantifind() Provider {
	return &quantifind{
		greenhouse{
			"Quantifind",
			"quantifind",
		},
	}
}
