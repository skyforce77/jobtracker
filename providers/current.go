package providers

type current struct {
	greenhouse
}

// NewCurrent returns a new provider
func NewCurrent() Provider {
	return &current{
		greenhouse{
			"Current",
			"current",
		},
	}
}
