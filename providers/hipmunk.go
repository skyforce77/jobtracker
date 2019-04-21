package providers

type hipmunk struct {
	greenhouse
}

// NewHipmunk returns a new provider
func NewHipmunk() Provider {
	return &hipmunk{
		greenhouse{
			"Hipmunk",
			"hipmunk",
		},
	}
}
