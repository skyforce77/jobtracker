package providers

type pocket struct {
	greenhouse
}

// NewPocket returns a new provider
func NewPocket() Provider {
	return &pocket{
		greenhouse{
			"Pocket",
			"pocket",
		},
	}
}
