package providers

type pocketgems struct {
	greenhouse
}

// NewPocketGems returns a new provider
func NewPocketGems() Provider {
	return &pocketgems{
		greenhouse{
			"PocketGems",
			"pocketgems",
		},
	}
}
