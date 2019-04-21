package providers

type lantern struct {
	greenhouse
}

// NewLantern returns a new provider
func NewLantern() Provider {
	return &lantern{
		greenhouse{
			"Lantern",
			"lantern",
		},
	}
}
