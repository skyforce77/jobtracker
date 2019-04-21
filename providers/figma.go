package providers

type figma struct {
	greenhouse
}

// NewFigma returns a new provider
func NewFigma() Provider {
	return &figma{
		greenhouse{
			"Figma",
			"figma",
		},
	}
}
