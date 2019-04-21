package providers

type joinhandshake struct {
	greenhouse
}

// NewHandshake returns a new provider
func NewHandshake() Provider {
	return &joinhandshake{
		greenhouse{
			"Handshake",
			"joinhandshake",
		},
	}
}
