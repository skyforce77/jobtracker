package providers

type justworks struct {
	greenhouse
}

// NewJustworks returns a new provider
func NewJustworks() Provider {
	return &justworks{
		greenhouse{
			"Justworks",
			"justworks",
		},
	}
}
