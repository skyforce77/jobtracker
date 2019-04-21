package providers

type improbable struct {
	greenhouse
}

// NewImprobable returns a new provider
func NewImprobable() Provider {
	return &improbable{
		greenhouse{
			"Improbable",
			"improbable",
		},
	}
}
