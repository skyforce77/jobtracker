package providers

type simplefinance struct {
	greenhouse
}

// NewSimpleFinance returns a new provider
func NewSimpleFinance() Provider {
	return &simplefinance{
		greenhouse{
			"SimpleFinance",
			"simplefinance",
		},
	}
}
