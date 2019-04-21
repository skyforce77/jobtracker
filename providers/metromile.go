package providers

type metromile struct {
	greenhouse
}

// NewMetromile returns a new provider
func NewMetromile() Provider {
	return &metromile{
		greenhouse{
			"Metromile",
			"metromile",
		},
	}
}
