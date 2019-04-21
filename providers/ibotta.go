package providers

type ibotta struct {
	greenhouse
}

// NewIbotta returns a new provider
func NewIbotta() Provider {
	return &ibotta{
		greenhouse{
			"Ibotta",
			"ibotta",
		},
	}
}
