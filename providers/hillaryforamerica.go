package providers

type hillaryforamerica struct {
	greenhouse
}

// NewHillaryForAmerica returns a new provider
func NewHillaryForAmerica() Provider {
	return &hillaryforamerica{
		greenhouse{
			"HillaryForAmerica",
			"hillaryforamerica",
		},
	}
}
