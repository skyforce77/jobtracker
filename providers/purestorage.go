package providers

type purestorage struct {
	greenhouse
}

// NewPureStorage returns a new provider
func NewPureStorage() Provider {
	return &purestorage{
		greenhouse{
			"PureStorage",
			"purestorage",
		},
	}
}
