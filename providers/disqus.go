package providers

type disqus struct {
	greenhouse
}

// NewDisqus returns a new provider
func NewDisqus() Provider {
	return &disqus{
		greenhouse{
			"Disqus",
			"disqus",
		},
	}
}
