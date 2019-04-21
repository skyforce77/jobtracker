package providers

type buzzfeed struct {
	greenhouse
}

// NewBuzzfeed returns a new provider
func NewBuzzfeed() Provider {
	return &buzzfeed{
		greenhouse{
			"Buzzfeed",
			"buzzfeed",
		},
	}
}
