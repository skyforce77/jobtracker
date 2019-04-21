package providers

type beekeeper struct {
	greenhouse
}

// NewBeekeeper returns a new provider
func NewBeekeeper() Provider {
	return &beekeeper{
		greenhouse{
			"Beekeeper",
			"beekeeper",
		},
	}
}
