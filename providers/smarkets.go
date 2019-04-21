package providers

type smarkets struct {
	greenhouse
}

// NewSmarkets returns a new provider
func NewSmarkets() Provider {
	return &smarkets{
		greenhouse{
			"Smarkets",
			"smarkets",
		},
	}
}
