package providers

type axiomzen struct {
	greenhouse
}

// NewAxiomZen returns a new provider
func NewAxiomZen() Provider {
	return &axiomzen{
		greenhouse{
			"AxiomZen",
			"axiomzen",
		},
	}
}
