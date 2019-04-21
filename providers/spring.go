package providers

type spring struct {
	greenhouse
}

// NewSpring returns a new provider
func NewSpring() Provider {
	return &spring{
		greenhouse{
			"Spring",
			"spring",
		},
	}
}
