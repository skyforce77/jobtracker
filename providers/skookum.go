package providers

type skookum struct {
	greenhouse
}

// NewSkookum returns a new provider
func NewSkookum() Provider {
	return &skookum{
		greenhouse{
			"Skookum",
			"skookum",
		},
	}
}
