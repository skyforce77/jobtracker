package providers

type carvana struct {
	greenhouse
}

// NewCarvana returns a new provider
func NewCarvana() Provider {
	return &carvana{
		greenhouse{
			"Carvana",
			"carvana",
		},
	}
}
