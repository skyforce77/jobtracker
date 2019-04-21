package providers

type cartodb struct {
	greenhouse
}

// NewCartoDB returns a new provider
func NewCartoDB() Provider {
	return &cartodb{
		greenhouse{
			"CartoDB",
			"cartodb",
		},
	}
}
