package providers

type digitalocean98 struct {
	greenhouse
}

// NewDigitalOcean returns a new provider
func NewDigitalOcean() Provider {
	return &digitalocean98{
		greenhouse{
			"DigitalOcean",
			"digitalocean98",
		},
	}
}
