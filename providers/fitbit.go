package providers

type fitbit92 struct {
	greenhouse
}

// NewFitbit returns a new provider
func NewFitbit() Provider {
	return &fitbit92{
		greenhouse{
			"Fitbit",
			"fitbit92",
		},
	}
}
