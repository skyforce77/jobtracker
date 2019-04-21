package providers

type rapid7 struct {
	greenhouse
}

// NewRapid7 returns a new provider
func NewRapid7() Provider {
	return &rapid7{
		greenhouse{
			"Rapid7",
			"rapid7",
		},
	}
}
