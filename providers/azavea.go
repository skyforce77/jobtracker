package providers

type azavea struct {
	greenhouse
}

// NewAzavea returns a new provider
func NewAzavea() Provider {
	return &azavea{
		greenhouse{
			"Azavea",
			"azavea",
		},
	}
}
