package providers

type bonobos struct {
	greenhouse
}

// NewBonobos returns a new provider
func NewBonobos() Provider {
	return &bonobos{
		greenhouse{
			"Bonobos",
			"bonobos",
		},
	}
}
