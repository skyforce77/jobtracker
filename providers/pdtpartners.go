package providers

type pdtpartners struct {
	greenhouse
}

// NewPDTPartners returns a new provider
func NewPDTPartners() Provider {
	return &pdtpartners{
		greenhouse{
			"PDTPartners",
			"pdtpartners",
		},
	}
}
