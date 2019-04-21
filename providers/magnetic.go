package providers

type magnetic struct {
	greenhouse
}

// NewMagnetic returns a new provider
func NewMagnetic() Provider {
	return &magnetic{
		greenhouse{
			"Magnetic",
			"magnetic",
		},
	}
}
