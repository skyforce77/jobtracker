package providers

type konekt struct {
	greenhouse
}

// NewKonekt returns a new provider
func NewKonekt() Provider {
	return &konekt{
		greenhouse{
			"Konekt",
			"konekt",
		},
	}
}
