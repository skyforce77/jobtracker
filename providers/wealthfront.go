package providers

type wealthfront struct {
	greenhouse
}

// NewWealthfront returns a new provider
func NewWealthfront() Provider {
	return &wealthfront{
		greenhouse{
			"Wealthfront",
			"wealthfront",
		},
	}
}
