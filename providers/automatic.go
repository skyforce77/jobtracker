package providers

type automatic struct {
	greenhouse
}

// NewAutomatic returns a new provider
func NewAutomatic() Provider {
	return &automatic{
		greenhouse{
			"Automatic",
			"automatic",
		},
	}
}
