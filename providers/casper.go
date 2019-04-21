package providers

type casper struct {
	greenhouse
}

// NewCasper returns a new provider
func NewCasper() Provider {
	return &casper{
		greenhouse{
			"Casper",
			"casper",
		},
	}
}
