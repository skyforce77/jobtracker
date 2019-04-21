package providers

type enernoc struct {
	greenhouse
}

// NewEnerNOC returns a new provider
func NewEnerNOC() Provider {
	return &enernoc{
		greenhouse{
			"EnerNOC",
			"enernoc",
		},
	}
}
