package providers

type pager struct {
	greenhouse
}

// NewPager returns a new provider
func NewPager() Provider {
	return &pager{
		greenhouse{
			"Pager",
			"pager",
		},
	}
}
