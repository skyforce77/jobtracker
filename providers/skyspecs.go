package providers

type skyspecs struct {
	greenhouse
}

// NewSkySpecs returns a new provider
func NewSkySpecs() Provider {
	return &skyspecs{
		greenhouse{
			"SkySpecs",
			"skyspecs",
		},
	}
}
