package providers

type placemeter struct {
	greenhouse
}

// NewPlacemeter returns a new provider
func NewPlacemeter() Provider {
	return &placemeter{
		greenhouse{
			"Placemeter",
			"placemeter",
		},
	}
}
