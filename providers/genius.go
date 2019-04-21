package providers

type genius struct {
	greenhouse
}

// NewGenius returns a new provider
func NewGenius() Provider {
	return &genius{
		greenhouse{
			"Genius",
			"genius",
		},
	}
}
