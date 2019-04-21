package providers

type unity3d struct {
	greenhouse
}

// NewUnity returns a new provider
func NewUnity() Provider {
	return &unity3d{
		greenhouse{
			"Unity",
			"unity3d",
		},
	}
}
