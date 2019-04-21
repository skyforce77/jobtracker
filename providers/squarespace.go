package providers

type squarespace struct {
	greenhouse
}

// NewSquarespace returns a new provider
func NewSquarespace() Provider {
	return &squarespace{
		greenhouse{
			"Squarespace",
			"squarespace",
		},
	}
}
