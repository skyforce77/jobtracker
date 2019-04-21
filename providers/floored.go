package providers

type floored struct {
	greenhouse
}

// NewFloored returns a new provider
func NewFloored() Provider {
	return &floored{
		greenhouse{
			"Floored",
			"floored",
		},
	}
}
