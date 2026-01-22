package providers

type figma struct {
	greenhouse
}

// NewFigma returns a new provider for Figma jobs
func NewFigma() Provider {
	return &figma{
		greenhouse{
			"Figma",
			"figma",
		},
	}
}
