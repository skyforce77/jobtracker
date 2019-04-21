package providers

type curse struct {
	greenhouse
}

// NewCurse returns a new provider
func NewCurse() Provider {
	return &curse{
		greenhouse{
			"Curse",
			"curse",
		},
	}
}
