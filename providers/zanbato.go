package providers

type zanbato struct {
	greenhouse
}

// NewZanbato returns a new provider
func NewZanbato() Provider {
	return &zanbato{
		greenhouse{
			"Zanbato",
			"zanbato",
		},
	}
}
