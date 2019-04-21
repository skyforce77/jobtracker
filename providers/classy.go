package providers

type classy struct {
	greenhouse
}

// NewClassy returns a new provider
func NewClassy() Provider {
	return &classy{
		greenhouse{
			"Classy",
			"classy",
		},
	}
}
