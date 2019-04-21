package providers

type massdrop struct {
	greenhouse
}

// NewMassdrop returns a new provider
func NewMassdrop() Provider {
	return &massdrop{
		greenhouse{
			"Massdrop",
			"massdrop",
		},
	}
}
