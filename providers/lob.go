package providers

type lob struct {
	greenhouse
}

// NewLob returns a new provider
func NewLob() Provider {
	return &lob{
		greenhouse{
			"Lob",
			"lob",
		},
	}
}
