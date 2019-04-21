package providers

type oseberg struct {
	greenhouse
}

// NewOseberg returns a new provider
func NewOseberg() Provider {
	return &oseberg{
		greenhouse{
			"Oseberg",
			"oseberg",
		},
	}
}
