package providers

type tumblr struct {
	greenhouse
}

// NewTumblr returns a new provider
func NewTumblr() Provider {
	return &tumblr{
		greenhouse{
			"Tumblr",
			"tumblr",
		},
	}
}
