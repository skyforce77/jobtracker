package providers

type signpost struct {
	greenhouse
}

// NewSignpost returns a new provider
func NewSignpost() Provider {
	return &signpost{
		greenhouse{
			"Signpost",
			"signpost",
		},
	}
}
