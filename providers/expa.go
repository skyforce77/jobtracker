package providers

type expa struct {
	greenhouse
}

// NewExpa returns a new provider
func NewExpa() Provider {
	return &expa{
		greenhouse{
			"Expa",
			"expa",
		},
	}
}
