package providers

type vimeo struct {
	greenhouse
}

// NewVimeo returns a new provider
func NewVimeo() Provider {
	return &vimeo{
		greenhouse{
			"Vimeo",
			"vimeo",
		},
	}
}
