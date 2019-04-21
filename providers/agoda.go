package providers

type agoda struct {
	greenhouse
}

// NewAgoda returns a new provider
func NewAgoda() Provider {
	return &agoda{
		greenhouse{
			"Agoda",
			"agoda",
		},
	}
}
