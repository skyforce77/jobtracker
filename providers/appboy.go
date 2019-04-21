package providers

type appboy struct {
	greenhouse
}

// NewAppboy returns a new provider
func NewAppboy() Provider {
	return &appboy{
		greenhouse{
			"Appboy",
			"appboy",
		},
	}
}
