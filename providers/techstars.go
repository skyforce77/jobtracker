package providers

type techstars struct {
	greenhouse
}

// NewTechstars returns a new provider
func NewTechstars() Provider {
	return &techstars{
		greenhouse{
			"Techstars",
			"techstars",
		},
	}
}
