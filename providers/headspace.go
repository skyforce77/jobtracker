package providers

type headspace struct {
	greenhouse
}

// NewHeadspace returns a new provider
func NewHeadspace() Provider {
	return &headspace{
		greenhouse{
			"Headspace",
			"headspace",
		},
	}
}
