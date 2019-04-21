package providers

type popsugar struct {
	greenhouse
}

// NewPopSugar returns a new provider
func NewPopSugar() Provider {
	return &popsugar{
		greenhouse{
			"PopSugar",
			"popsugar",
		},
	}
}
