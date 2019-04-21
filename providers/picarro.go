package providers

type picarro struct {
	greenhouse
}

// NewPicarro returns a new provider
func NewPicarro() Provider {
	return &picarro{
		greenhouse{
			"Picarro",
			"picarro",
		},
	}
}
