package providers

type peek struct {
	greenhouse
}

// NewPeek returns a new provider
func NewPeek() Provider {
	return &peek{
		greenhouse{
			"Peek",
			"peek",
		},
	}
}
