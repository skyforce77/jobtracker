package providers

type twg struct {
	greenhouse
}

// NewTheWorkingGroup returns a new provider
func NewTheWorkingGroup() Provider {
	return &twg{
		greenhouse{
			"TheWorkingGroup",
			"twg",
		},
	}
}
