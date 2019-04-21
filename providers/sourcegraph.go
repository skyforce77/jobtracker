package providers

type sourcegraph struct {
	greenhouse
}

// NewSourcegraph returns a new provider
func NewSourcegraph() Provider {
	return &sourcegraph{
		greenhouse{
			"Sourcegraph",
			"sourcegraph",
		},
	}
}
