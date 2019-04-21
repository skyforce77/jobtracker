package providers

type zype struct {
	greenhouse
}

// NewZype returns a new provider
func NewZype() Provider {
	return &zype{
		greenhouse{
			"Zype",
			"zype",
		},
	}
}
