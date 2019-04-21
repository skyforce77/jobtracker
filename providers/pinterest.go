package providers

type pinterest struct {
	greenhouse
}

// NewPinterest returns a new provider
func NewPinterest() Provider {
	return &pinterest{
		greenhouse{
			"Pinterest",
			"pinterest",
		},
	}
}
