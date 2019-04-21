package providers

type github struct {
	greenhouse
}

// NewGithub returns a new provider
func NewGithub() Provider {
	return &github{
		greenhouse{
			"Github",
			"github",
		},
	}
}
